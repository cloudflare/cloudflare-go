// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustGatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustGatewayRuleService]
// method instead.
type ZeroTrustGatewayRuleService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayRuleService(opts ...option.RequestOption) (r *ZeroTrustGatewayRuleService) {
	r = &ZeroTrustGatewayRuleService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *ZeroTrustGatewayRuleService) New(ctx context.Context, params ZeroTrustGatewayRuleNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *ZeroTrustGatewayRuleService) Update(ctx context.Context, ruleID string, params ZeroTrustGatewayRuleUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *ZeroTrustGatewayRuleService) List(ctx context.Context, query ZeroTrustGatewayRuleListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust Gateway rule.
func (r *ZeroTrustGatewayRuleService) Delete(ctx context.Context, ruleID string, body ZeroTrustGatewayRuleDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *ZeroTrustGatewayRuleService) Get(ctx context.Context, ruleID string, query ZeroTrustGatewayRuleGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
)

// The protocol or layer to use.
type ZeroTrustGatewayRulesFilter string

const (
	ZeroTrustGatewayRulesFilterHTTP   ZeroTrustGatewayRulesFilter = "http"
	ZeroTrustGatewayRulesFilterDNS    ZeroTrustGatewayRulesFilter = "dns"
	ZeroTrustGatewayRulesFilterL4     ZeroTrustGatewayRulesFilter = "l4"
	ZeroTrustGatewayRulesFilterEgress ZeroTrustGatewayRulesFilter = "egress"
)

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
	// route to the address closest to their origin.
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
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert ZeroTrustGatewayRulesRuleSettingsUntrustedCert `json:"untrusted_cert"`
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
	UntrustedCert                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
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

type ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
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

type ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
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

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type ZeroTrustGatewayRulesRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action ZeroTrustGatewayRulesRuleSettingsUntrustedCertAction `json:"action"`
	JSON   zeroTrustGatewayRulesRuleSettingsUntrustedCertJSON   `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsUntrustedCertJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRulesRuleSettingsUntrustedCert]
type zeroTrustGatewayRulesRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type ZeroTrustGatewayRulesRuleSettingsUntrustedCertAction string

const (
	ZeroTrustGatewayRulesRuleSettingsUntrustedCertActionPassThrough ZeroTrustGatewayRulesRuleSettingsUntrustedCertAction = "pass_through"
	ZeroTrustGatewayRulesRuleSettingsUntrustedCertActionBlock       ZeroTrustGatewayRulesRuleSettingsUntrustedCertAction = "block"
	ZeroTrustGatewayRulesRuleSettingsUntrustedCertActionError       ZeroTrustGatewayRulesRuleSettingsUntrustedCertAction = "error"
)

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

// Union satisfied by [ZeroTrustGatewayRuleDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustGatewayRuleDeleteResponse interface {
	ImplementsZeroTrustGatewayRuleDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayRuleDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustGatewayRuleNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[ZeroTrustGatewayRuleNewParamsAction] `json:"action,required"`
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
	Filters param.Field[[]ZeroTrustGatewayRuleNewParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[ZeroTrustGatewayRuleNewParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ZeroTrustGatewayRuleNewParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r ZeroTrustGatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type ZeroTrustGatewayRuleNewParamsAction string

const (
	ZeroTrustGatewayRuleNewParamsActionOn           ZeroTrustGatewayRuleNewParamsAction = "on"
	ZeroTrustGatewayRuleNewParamsActionOff          ZeroTrustGatewayRuleNewParamsAction = "off"
	ZeroTrustGatewayRuleNewParamsActionAllow        ZeroTrustGatewayRuleNewParamsAction = "allow"
	ZeroTrustGatewayRuleNewParamsActionBlock        ZeroTrustGatewayRuleNewParamsAction = "block"
	ZeroTrustGatewayRuleNewParamsActionScan         ZeroTrustGatewayRuleNewParamsAction = "scan"
	ZeroTrustGatewayRuleNewParamsActionNoscan       ZeroTrustGatewayRuleNewParamsAction = "noscan"
	ZeroTrustGatewayRuleNewParamsActionSafesearch   ZeroTrustGatewayRuleNewParamsAction = "safesearch"
	ZeroTrustGatewayRuleNewParamsActionYtrestricted ZeroTrustGatewayRuleNewParamsAction = "ytrestricted"
	ZeroTrustGatewayRuleNewParamsActionIsolate      ZeroTrustGatewayRuleNewParamsAction = "isolate"
	ZeroTrustGatewayRuleNewParamsActionNoisolate    ZeroTrustGatewayRuleNewParamsAction = "noisolate"
	ZeroTrustGatewayRuleNewParamsActionOverride     ZeroTrustGatewayRuleNewParamsAction = "override"
	ZeroTrustGatewayRuleNewParamsActionL4Override   ZeroTrustGatewayRuleNewParamsAction = "l4_override"
	ZeroTrustGatewayRuleNewParamsActionEgress       ZeroTrustGatewayRuleNewParamsAction = "egress"
	ZeroTrustGatewayRuleNewParamsActionAuditSSH     ZeroTrustGatewayRuleNewParamsAction = "audit_ssh"
)

// The protocol or layer to use.
type ZeroTrustGatewayRuleNewParamsFilter string

const (
	ZeroTrustGatewayRuleNewParamsFilterHTTP   ZeroTrustGatewayRuleNewParamsFilter = "http"
	ZeroTrustGatewayRuleNewParamsFilterDNS    ZeroTrustGatewayRuleNewParamsFilter = "dns"
	ZeroTrustGatewayRuleNewParamsFilterL4     ZeroTrustGatewayRuleNewParamsFilter = "l4"
	ZeroTrustGatewayRuleNewParamsFilterEgress ZeroTrustGatewayRuleNewParamsFilter = "egress"
)

// Additional settings that modify the rule's action.
type ZeroTrustGatewayRuleNewParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type ZeroTrustGatewayRuleNewParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type ZeroTrustGatewayRuleNewParamsRuleSettingsBisoAdminControls struct {
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

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type ZeroTrustGatewayRuleNewParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsDNSResolversIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type ZeroTrustGatewayRuleNewParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type ZeroTrustGatewayRuleNewParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type ZeroTrustGatewayRuleNewParamsRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type ZeroTrustGatewayRuleNewParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertAction string

const (
	ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertActionPassThrough ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertAction = "pass_through"
	ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertActionBlock       ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertAction = "block"
	ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertActionError       ZeroTrustGatewayRuleNewParamsRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ZeroTrustGatewayRuleNewParamsSchedule struct {
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

func (r ZeroTrustGatewayRuleNewParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleNewResponseEnvelope struct {
	Errors   []ZeroTrustGatewayRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayRules                             `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayRuleNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayRuleNewResponseEnvelope]
type zeroTrustGatewayRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustGatewayRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRuleNewResponseEnvelopeErrors]
type zeroTrustGatewayRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustGatewayRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleNewResponseEnvelopeMessages]
type zeroTrustGatewayRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayRuleNewResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayRuleNewResponseEnvelopeSuccessTrue ZeroTrustGatewayRuleNewResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayRuleUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[ZeroTrustGatewayRuleUpdateParamsAction] `json:"action,required"`
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
	Filters param.Field[[]ZeroTrustGatewayRuleUpdateParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ZeroTrustGatewayRuleUpdateParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r ZeroTrustGatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type ZeroTrustGatewayRuleUpdateParamsAction string

const (
	ZeroTrustGatewayRuleUpdateParamsActionOn           ZeroTrustGatewayRuleUpdateParamsAction = "on"
	ZeroTrustGatewayRuleUpdateParamsActionOff          ZeroTrustGatewayRuleUpdateParamsAction = "off"
	ZeroTrustGatewayRuleUpdateParamsActionAllow        ZeroTrustGatewayRuleUpdateParamsAction = "allow"
	ZeroTrustGatewayRuleUpdateParamsActionBlock        ZeroTrustGatewayRuleUpdateParamsAction = "block"
	ZeroTrustGatewayRuleUpdateParamsActionScan         ZeroTrustGatewayRuleUpdateParamsAction = "scan"
	ZeroTrustGatewayRuleUpdateParamsActionNoscan       ZeroTrustGatewayRuleUpdateParamsAction = "noscan"
	ZeroTrustGatewayRuleUpdateParamsActionSafesearch   ZeroTrustGatewayRuleUpdateParamsAction = "safesearch"
	ZeroTrustGatewayRuleUpdateParamsActionYtrestricted ZeroTrustGatewayRuleUpdateParamsAction = "ytrestricted"
	ZeroTrustGatewayRuleUpdateParamsActionIsolate      ZeroTrustGatewayRuleUpdateParamsAction = "isolate"
	ZeroTrustGatewayRuleUpdateParamsActionNoisolate    ZeroTrustGatewayRuleUpdateParamsAction = "noisolate"
	ZeroTrustGatewayRuleUpdateParamsActionOverride     ZeroTrustGatewayRuleUpdateParamsAction = "override"
	ZeroTrustGatewayRuleUpdateParamsActionL4Override   ZeroTrustGatewayRuleUpdateParamsAction = "l4_override"
	ZeroTrustGatewayRuleUpdateParamsActionEgress       ZeroTrustGatewayRuleUpdateParamsAction = "egress"
	ZeroTrustGatewayRuleUpdateParamsActionAuditSSH     ZeroTrustGatewayRuleUpdateParamsAction = "audit_ssh"
)

// The protocol or layer to use.
type ZeroTrustGatewayRuleUpdateParamsFilter string

const (
	ZeroTrustGatewayRuleUpdateParamsFilterHTTP   ZeroTrustGatewayRuleUpdateParamsFilter = "http"
	ZeroTrustGatewayRuleUpdateParamsFilterDNS    ZeroTrustGatewayRuleUpdateParamsFilter = "dns"
	ZeroTrustGatewayRuleUpdateParamsFilterL4     ZeroTrustGatewayRuleUpdateParamsFilter = "l4"
	ZeroTrustGatewayRuleUpdateParamsFilterEgress ZeroTrustGatewayRuleUpdateParamsFilter = "egress"
)

// Additional settings that modify the rule's action.
type ZeroTrustGatewayRuleUpdateParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsBisoAdminControls struct {
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

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction string

const (
	ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionPassThrough ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "pass_through"
	ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionBlock       ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "block"
	ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError       ZeroTrustGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ZeroTrustGatewayRuleUpdateParamsSchedule struct {
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

func (r ZeroTrustGatewayRuleUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayRuleUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayRules                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayRuleUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRuleUpdateResponseEnvelope]
type zeroTrustGatewayRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleUpdateResponseEnvelopeErrors]
type zeroTrustGatewayRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayRuleUpdateResponseEnvelopeMessages]
type zeroTrustGatewayRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayRuleUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayRuleUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayRuleListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayRuleListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayRules                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayRuleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayRuleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayRuleListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayRuleListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayRuleListResponseEnvelope]
type zeroTrustGatewayRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustGatewayRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayRuleListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleListResponseEnvelopeErrors]
type zeroTrustGatewayRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayRuleListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleListResponseEnvelopeMessages]
type zeroTrustGatewayRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayRuleListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayRuleListResponseEnvelopeSuccessTrue ZeroTrustGatewayRuleListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayRuleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zeroTrustGatewayRuleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayRuleListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayRuleListResponseEnvelopeResultInfo]
type zeroTrustGatewayRuleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayRuleDeleteResponseEnvelope struct {
	Errors   []ZeroTrustGatewayRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayRuleDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayRuleDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRuleDeleteResponseEnvelope]
type zeroTrustGatewayRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleDeleteResponseEnvelopeErrors]
type zeroTrustGatewayRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayRuleDeleteResponseEnvelopeMessages]
type zeroTrustGatewayRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayRuleDeleteResponseEnvelopeSuccessTrue ZeroTrustGatewayRuleDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayRuleGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayRuleGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayRules                             `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayRuleGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayRuleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayRuleGetResponseEnvelope]
type zeroTrustGatewayRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustGatewayRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRuleGetResponseEnvelopeErrors]
type zeroTrustGatewayRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayRuleGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustGatewayRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRuleGetResponseEnvelopeMessages]
type zeroTrustGatewayRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayRuleGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayRuleGetResponseEnvelopeSuccessTrue ZeroTrustGatewayRuleGetResponseEnvelopeSuccess = true
)

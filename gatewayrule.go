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

// Updates a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, accountID interface{}, ruleID string, body GatewayRuleUpdateParams, opts ...option.RequestOption) (res *GatewayRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, accountID interface{}, ruleID string, opts ...option.RequestOption) (res *GatewayRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, accountID interface{}, ruleID string, opts ...option.RequestOption) (res *GatewayRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *GatewayRuleService) ZeroTrustGatewayRulesNewZeroTrustGatewayRule(ctx context.Context, accountID interface{}, body GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams, opts ...option.RequestOption) (res *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) ZeroTrustGatewayRulesListZeroTrustGatewayRules(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayRuleUpdateResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleUpdateResponseAction `json:"action"`
	CreatedAt time.Time                       `json:"created_at" format:"date-time"`
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
	Filters []GatewayRuleUpdateResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleUpdateResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleUpdateResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                        `json:"traffic"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleUpdateResponseJSON `json:"-"`
}

// gatewayRuleUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponse]
type gatewayRuleUpdateResponseJSON struct {
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

func (r *GatewayRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleUpdateResponseAction string

const (
	GatewayRuleUpdateResponseActionOn           GatewayRuleUpdateResponseAction = "on"
	GatewayRuleUpdateResponseActionOff          GatewayRuleUpdateResponseAction = "off"
	GatewayRuleUpdateResponseActionAllow        GatewayRuleUpdateResponseAction = "allow"
	GatewayRuleUpdateResponseActionBlock        GatewayRuleUpdateResponseAction = "block"
	GatewayRuleUpdateResponseActionScan         GatewayRuleUpdateResponseAction = "scan"
	GatewayRuleUpdateResponseActionNoscan       GatewayRuleUpdateResponseAction = "noscan"
	GatewayRuleUpdateResponseActionSafesearch   GatewayRuleUpdateResponseAction = "safesearch"
	GatewayRuleUpdateResponseActionYtrestricted GatewayRuleUpdateResponseAction = "ytrestricted"
	GatewayRuleUpdateResponseActionIsolate      GatewayRuleUpdateResponseAction = "isolate"
	GatewayRuleUpdateResponseActionNoisolate    GatewayRuleUpdateResponseAction = "noisolate"
	GatewayRuleUpdateResponseActionOverride     GatewayRuleUpdateResponseAction = "override"
	GatewayRuleUpdateResponseActionL4Override   GatewayRuleUpdateResponseAction = "l4_override"
	GatewayRuleUpdateResponseActionEgress       GatewayRuleUpdateResponseAction = "egress"
	GatewayRuleUpdateResponseActionAuditSSH     GatewayRuleUpdateResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleUpdateResponseFilter string

const (
	GatewayRuleUpdateResponseFilterHTTP   GatewayRuleUpdateResponseFilter = "http"
	GatewayRuleUpdateResponseFilterDNS    GatewayRuleUpdateResponseFilter = "dns"
	GatewayRuleUpdateResponseFilterL4     GatewayRuleUpdateResponseFilter = "l4"
	GatewayRuleUpdateResponseFilterEgress GatewayRuleUpdateResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleUpdateResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleUpdateResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleUpdateResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleUpdateResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleUpdateResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleUpdateResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleUpdateResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleUpdateResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleUpdateResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleUpdateResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleUpdateResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsJSON contains the JSON metadata for the
// struct [GatewayRuleUpdateResponseRuleSettings]
type gatewayRuleUpdateResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleUpdateResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type GatewayRuleUpdateResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                              `json:"command_logging"`
	JSON           gatewayRuleUpdateResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsAuditSSHJSON contains the JSON metadata for
// the struct [GatewayRuleUpdateResponseRuleSettingsAuditSSH]
type gatewayRuleUpdateResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type GatewayRuleUpdateResponseRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                       `json:"du"`
	JSON gatewayRuleUpdateResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct [GatewayRuleUpdateResponseRuleSettingsBisoAdminControls]
type gatewayRuleUpdateResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type GatewayRuleUpdateResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                  `json:"enforce"`
	JSON    gatewayRuleUpdateResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsCheckSessionJSON contains the JSON metadata
// for the struct [GatewayRuleUpdateResponseRuleSettingsCheckSession]
type gatewayRuleUpdateResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleUpdateResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleUpdateResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsDNSResolversJSON contains the JSON metadata
// for the struct [GatewayRuleUpdateResponseRuleSettingsDNSResolvers]
type gatewayRuleUpdateResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                    `json:"vnet_id"`
	JSON   gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4JSON contains the JSON
// metadata for the struct [GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                    `json:"vnet_id"`
	JSON   gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6JSON contains the JSON
// metadata for the struct [GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleUpdateResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                          `json:"ipv6"`
	JSON gatewayRuleUpdateResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsEgressJSON contains the JSON metadata for
// the struct [GatewayRuleUpdateResponseRuleSettingsEgress]
type gatewayRuleUpdateResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleUpdateResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                               `json:"port"`
	JSON gatewayRuleUpdateResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsL4overrideJSON contains the JSON metadata
// for the struct [GatewayRuleUpdateResponseRuleSettingsL4override]
type gatewayRuleUpdateResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleUpdateResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                        `json:"support_url"`
	JSON       gatewayRuleUpdateResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsNotificationSettingsJSON contains the JSON
// metadata for the struct
// [GatewayRuleUpdateResponseRuleSettingsNotificationSettings]
type gatewayRuleUpdateResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type GatewayRuleUpdateResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                `json:"enabled"`
	JSON    gatewayRuleUpdateResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsPayloadLogJSON contains the JSON metadata
// for the struct [GatewayRuleUpdateResponseRuleSettingsPayloadLog]
type gatewayRuleUpdateResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleUpdateResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleUpdateResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleUpdateResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleUpdateResponseRuleSettingsUntrustedCertJSON contains the JSON
// metadata for the struct [GatewayRuleUpdateResponseRuleSettingsUntrustedCert]
type gatewayRuleUpdateResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleUpdateResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleUpdateResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleUpdateResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleUpdateResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleUpdateResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleUpdateResponseRuleSettingsUntrustedCertActionError       GatewayRuleUpdateResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleUpdateResponseSchedule struct {
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
	Wed  string                                `json:"wed"`
	JSON gatewayRuleUpdateResponseScheduleJSON `json:"-"`
}

// gatewayRuleUpdateResponseScheduleJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponseSchedule]
type gatewayRuleUpdateResponseScheduleJSON struct {
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

func (r *GatewayRuleUpdateResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [GatewayRuleDeleteResponseUnknown] or [shared.UnionString].
type GatewayRuleDeleteResponse interface {
	ImplementsGatewayRuleDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayRuleDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayRuleGetResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleGetResponseAction `json:"action"`
	CreatedAt time.Time                    `json:"created_at" format:"date-time"`
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
	Filters []GatewayRuleGetResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleGetResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleGetResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                     `json:"traffic"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleGetResponseJSON `json:"-"`
}

// gatewayRuleGetResponseJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponse]
type gatewayRuleGetResponseJSON struct {
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

func (r *GatewayRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleGetResponseAction string

const (
	GatewayRuleGetResponseActionOn           GatewayRuleGetResponseAction = "on"
	GatewayRuleGetResponseActionOff          GatewayRuleGetResponseAction = "off"
	GatewayRuleGetResponseActionAllow        GatewayRuleGetResponseAction = "allow"
	GatewayRuleGetResponseActionBlock        GatewayRuleGetResponseAction = "block"
	GatewayRuleGetResponseActionScan         GatewayRuleGetResponseAction = "scan"
	GatewayRuleGetResponseActionNoscan       GatewayRuleGetResponseAction = "noscan"
	GatewayRuleGetResponseActionSafesearch   GatewayRuleGetResponseAction = "safesearch"
	GatewayRuleGetResponseActionYtrestricted GatewayRuleGetResponseAction = "ytrestricted"
	GatewayRuleGetResponseActionIsolate      GatewayRuleGetResponseAction = "isolate"
	GatewayRuleGetResponseActionNoisolate    GatewayRuleGetResponseAction = "noisolate"
	GatewayRuleGetResponseActionOverride     GatewayRuleGetResponseAction = "override"
	GatewayRuleGetResponseActionL4Override   GatewayRuleGetResponseAction = "l4_override"
	GatewayRuleGetResponseActionEgress       GatewayRuleGetResponseAction = "egress"
	GatewayRuleGetResponseActionAuditSSH     GatewayRuleGetResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleGetResponseFilter string

const (
	GatewayRuleGetResponseFilterHTTP   GatewayRuleGetResponseFilter = "http"
	GatewayRuleGetResponseFilterDNS    GatewayRuleGetResponseFilter = "dns"
	GatewayRuleGetResponseFilterL4     GatewayRuleGetResponseFilter = "l4"
	GatewayRuleGetResponseFilterEgress GatewayRuleGetResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleGetResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleGetResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleGetResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleGetResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleGetResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleGetResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleGetResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleGetResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleGetResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleGetResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleGetResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseRuleSettings]
type gatewayRuleGetResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleGetResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type GatewayRuleGetResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                           `json:"command_logging"`
	JSON           gatewayRuleGetResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsAuditSSHJSON contains the JSON metadata for
// the struct [GatewayRuleGetResponseRuleSettingsAuditSSH]
type gatewayRuleGetResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type GatewayRuleGetResponseRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                    `json:"du"`
	JSON gatewayRuleGetResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct [GatewayRuleGetResponseRuleSettingsBisoAdminControls]
type gatewayRuleGetResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type GatewayRuleGetResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                               `json:"enforce"`
	JSON    gatewayRuleGetResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsCheckSessionJSON contains the JSON metadata
// for the struct [GatewayRuleGetResponseRuleSettingsCheckSession]
type gatewayRuleGetResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleGetResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleGetResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleGetResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleGetResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsDNSResolversJSON contains the JSON metadata
// for the struct [GatewayRuleGetResponseRuleSettingsDNSResolvers]
type gatewayRuleGetResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleGetResponseRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                 `json:"vnet_id"`
	JSON   gatewayRuleGetResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsDNSResolversIPV4JSON contains the JSON
// metadata for the struct [GatewayRuleGetResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleGetResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleGetResponseRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                 `json:"vnet_id"`
	JSON   gatewayRuleGetResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsDNSResolversIPV6JSON contains the JSON
// metadata for the struct [GatewayRuleGetResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleGetResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleGetResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                       `json:"ipv6"`
	JSON gatewayRuleGetResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsEgressJSON contains the JSON metadata for the
// struct [GatewayRuleGetResponseRuleSettingsEgress]
type gatewayRuleGetResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleGetResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                            `json:"port"`
	JSON gatewayRuleGetResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsL4overrideJSON contains the JSON metadata for
// the struct [GatewayRuleGetResponseRuleSettingsL4override]
type gatewayRuleGetResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleGetResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                     `json:"support_url"`
	JSON       gatewayRuleGetResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsNotificationSettingsJSON contains the JSON
// metadata for the struct [GatewayRuleGetResponseRuleSettingsNotificationSettings]
type gatewayRuleGetResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type GatewayRuleGetResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                             `json:"enabled"`
	JSON    gatewayRuleGetResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsPayloadLogJSON contains the JSON metadata for
// the struct [GatewayRuleGetResponseRuleSettingsPayloadLog]
type gatewayRuleGetResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleGetResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleGetResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleGetResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleGetResponseRuleSettingsUntrustedCertJSON contains the JSON metadata
// for the struct [GatewayRuleGetResponseRuleSettingsUntrustedCert]
type gatewayRuleGetResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleGetResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleGetResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleGetResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleGetResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleGetResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleGetResponseRuleSettingsUntrustedCertActionError       GatewayRuleGetResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleGetResponseSchedule struct {
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
	Wed  string                             `json:"wed"`
	JSON gatewayRuleGetResponseScheduleJSON `json:"-"`
}

// gatewayRuleGetResponseScheduleJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseSchedule]
type gatewayRuleGetResponseScheduleJSON struct {
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

func (r *GatewayRuleGetResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction `json:"action"`
	CreatedAt time.Time                                                             `json:"created_at" format:"date-time"`
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
	Filters []GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                                              `json:"traffic"`
	UpdatedAt time.Time                                                           `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON contains the
// JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionOn           GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "on"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionOff          GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "off"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionAllow        GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "allow"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionBlock        GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "block"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionScan         GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "scan"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionNoscan       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "noscan"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionSafesearch   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "safesearch"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionYtrestricted GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "ytrestricted"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionIsolate      GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "isolate"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionNoisolate    GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "noisolate"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionOverride     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "override"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionL4Override   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "l4_override"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionEgress       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "egress"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseActionAuditSSH     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilterHTTP   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter = "http"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilterDNS    GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter = "dns"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilterL4     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter = "l4"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilterEgress GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettings]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                                                    `json:"command_logging"`
	JSON           gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSHJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSH]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                                             `json:"du"`
	JSON gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControlsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControls]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                                                        `json:"enforce"`
	JSON    gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSessionJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSession]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolvers]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                          `json:"vnet_id"`
	JSON   gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4JSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                          `json:"vnet_id"`
	JSON   gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6JSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                                                                `json:"ipv6"`
	JSON gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgressJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgress]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                                                     `json:"port"`
	JSON gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4overrideJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4override]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                              `json:"support_url"`
	JSON       gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettings]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                                                      `json:"enabled"`
	JSON    gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLogJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLog]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCert]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertActionError       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSchedule struct {
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
	Wed  string                                                                      `json:"wed"`
	JSON gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseScheduleJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseScheduleJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSchedule]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseScheduleJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction `json:"action"`
	CreatedAt time.Time                                                               `json:"created_at" format:"date-time"`
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
	Filters []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                                                `json:"traffic"`
	UpdatedAt time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON contains
// the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction string

const (
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionOn           GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "on"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionOff          GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "off"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionAllow        GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "allow"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionBlock        GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "block"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionScan         GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "scan"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionNoscan       GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "noscan"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionSafesearch   GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "safesearch"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionYtrestricted GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "ytrestricted"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionIsolate      GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "isolate"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionNoisolate    GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "noisolate"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionOverride     GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "override"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionL4Override   GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "l4_override"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionEgress       GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "egress"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseActionAuditSSH     GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter string

const (
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilterHTTP   GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter = "http"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilterDNS    GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter = "dns"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilterL4     GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter = "l4"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilterEgress GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettings]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                                                      `json:"command_logging"`
	JSON           gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSHJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSH]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                                               `json:"du"`
	JSON gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControlsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControls]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                                                          `json:"enforce"`
	JSON    gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSessionJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSession]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolvers]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                            `json:"vnet_id"`
	JSON   gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4JSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                            `json:"vnet_id"`
	JSON   gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6JSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                                                                  `json:"ipv6"`
	JSON gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgressJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgress]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                                                       `json:"port"`
	JSON gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4overrideJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4override]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                `json:"support_url"`
	JSON       gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettings]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                                                        `json:"enabled"`
	JSON    gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLogJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLog]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCert]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertActionError       GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSchedule struct {
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
	Wed  string                                                                        `json:"wed"`
	JSON gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseScheduleJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseScheduleJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSchedule]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseScheduleJSON struct {
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

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleUpdateParams struct {
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
)

// The protocol or layer to use.
type GatewayRuleUpdateParamsFilter string

const (
	GatewayRuleUpdateParamsFilterHTTP   GatewayRuleUpdateParamsFilter = "http"
	GatewayRuleUpdateParamsFilterDNS    GatewayRuleUpdateParamsFilter = "dns"
	GatewayRuleUpdateParamsFilterL4     GatewayRuleUpdateParamsFilter = "l4"
	GatewayRuleUpdateParamsFilterEgress GatewayRuleUpdateParamsFilter = "egress"
)

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
	// route to the address closest to their origin.
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
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[GatewayRuleUpdateParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
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
// route to the address closest to their origin.
type GatewayRuleUpdateParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r GatewayRuleUpdateParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4 struct {
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

func (r GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6 struct {
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
type GatewayRuleUpdateParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[GatewayRuleUpdateParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r GatewayRuleUpdateParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleUpdateParamsRuleSettingsUntrustedCertAction string

const (
	GatewayRuleUpdateParamsRuleSettingsUntrustedCertActionPassThrough GatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleUpdateParamsRuleSettingsUntrustedCertActionBlock       GatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "block"
	GatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError       GatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "error"
)

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
	Errors   []GatewayRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayRuleUpdateResponse                   `json:"result,required"`
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

type GatewayRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayRuleUpdateResponseEnvelopeErrors]
type gatewayRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayRuleUpdateResponseEnvelopeMessages]
type gatewayRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	GatewayRuleUpdateResponseEnvelopeSuccessTrue GatewayRuleUpdateResponseEnvelopeSuccess = true
)

type GatewayRuleDeleteResponseEnvelope struct {
	Errors   []GatewayRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayRuleDeleteResponse                   `json:"result,required"`
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

type GatewayRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayRuleDeleteResponseEnvelopeErrors]
type gatewayRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayRuleDeleteResponseEnvelopeMessages]
type gatewayRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	GatewayRuleDeleteResponseEnvelopeSuccessTrue GatewayRuleDeleteResponseEnvelopeSuccess = true
)

type GatewayRuleGetResponseEnvelope struct {
	Errors   []GatewayRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayRuleGetResponse                   `json:"result,required"`
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

type GatewayRuleGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    gatewayRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayRuleGetResponseEnvelopeErrors]
type gatewayRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    gatewayRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayRuleGetResponseEnvelopeMessages]
type gatewayRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayRuleGetResponseEnvelopeSuccess bool

const (
	GatewayRuleGetResponseEnvelopeSuccessTrue GatewayRuleGetResponseEnvelopeSuccess = true
)

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams struct {
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction] `json:"action,required"`
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
	Filters param.Field[[]GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOn           GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "on"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOff          GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "off"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionAllow        GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "allow"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionBlock        GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "block"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionScan         GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "scan"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionNoscan       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "noscan"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionSafesearch   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "safesearch"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionYtrestricted GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "ytrestricted"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionIsolate      GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "isolate"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionNoisolate    GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "noisolate"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOverride     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "override"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionL4Override   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "l4_override"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionEgress       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "egress"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionAuditSSH     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterHTTP   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "http"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterDNS    GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "dns"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterL4     GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "l4"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterEgress GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls struct {
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

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV4 struct {
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

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV6 struct {
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

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction string

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionPassThrough GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionBlock       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "block"
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionError       GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule struct {
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

func (r GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelope struct {
	Errors   []GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelope]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrors]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessages]
type gatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeSuccess bool

const (
	GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeSuccessTrue GatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseEnvelopeSuccess = true
)

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelope struct {
	Errors   []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeJSON       `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelope]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrors]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessages]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeSuccess bool

const (
	GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeSuccessTrue GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeSuccess = true
)

type GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                 `json:"total_count"`
	JSON       gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfo]
type gatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

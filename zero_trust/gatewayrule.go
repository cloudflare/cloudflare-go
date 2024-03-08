// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *GatewayRuleService) New(ctx context.Context, params GatewayRuleNewParams, opts ...option.RequestOption) (res *GatewayRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, ruleID string, params GatewayRuleUpdateParams, opts ...option.RequestOption) (res *GatewayRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) List(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) (res *[]GatewayRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, ruleID string, body GatewayRuleDeleteParams, opts ...option.RequestOption) (res *GatewayRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, ruleID string, query GatewayRuleGetParams, opts ...option.RequestOption) (res *GatewayRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayRuleNewResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleNewResponseAction `json:"action"`
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
	Filters []GatewayRuleNewResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleNewResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleNewResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                     `json:"traffic"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleNewResponseJSON `json:"-"`
}

// gatewayRuleNewResponseJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponse]
type gatewayRuleNewResponseJSON struct {
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

func (r *GatewayRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleNewResponseAction string

const (
	GatewayRuleNewResponseActionOn           GatewayRuleNewResponseAction = "on"
	GatewayRuleNewResponseActionOff          GatewayRuleNewResponseAction = "off"
	GatewayRuleNewResponseActionAllow        GatewayRuleNewResponseAction = "allow"
	GatewayRuleNewResponseActionBlock        GatewayRuleNewResponseAction = "block"
	GatewayRuleNewResponseActionScan         GatewayRuleNewResponseAction = "scan"
	GatewayRuleNewResponseActionNoscan       GatewayRuleNewResponseAction = "noscan"
	GatewayRuleNewResponseActionSafesearch   GatewayRuleNewResponseAction = "safesearch"
	GatewayRuleNewResponseActionYtrestricted GatewayRuleNewResponseAction = "ytrestricted"
	GatewayRuleNewResponseActionIsolate      GatewayRuleNewResponseAction = "isolate"
	GatewayRuleNewResponseActionNoisolate    GatewayRuleNewResponseAction = "noisolate"
	GatewayRuleNewResponseActionOverride     GatewayRuleNewResponseAction = "override"
	GatewayRuleNewResponseActionL4Override   GatewayRuleNewResponseAction = "l4_override"
	GatewayRuleNewResponseActionEgress       GatewayRuleNewResponseAction = "egress"
	GatewayRuleNewResponseActionAuditSSH     GatewayRuleNewResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleNewResponseFilter string

const (
	GatewayRuleNewResponseFilterHTTP   GatewayRuleNewResponseFilter = "http"
	GatewayRuleNewResponseFilterDNS    GatewayRuleNewResponseFilter = "dns"
	GatewayRuleNewResponseFilterL4     GatewayRuleNewResponseFilter = "l4"
	GatewayRuleNewResponseFilterEgress GatewayRuleNewResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleNewResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleNewResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleNewResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleNewResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleNewResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleNewResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleNewResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleNewResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleNewResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleNewResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleNewResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseRuleSettings]
type gatewayRuleNewResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleNewResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type GatewayRuleNewResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                           `json:"command_logging"`
	JSON           gatewayRuleNewResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsAuditSSHJSON contains the JSON metadata for
// the struct [GatewayRuleNewResponseRuleSettingsAuditSSH]
type gatewayRuleNewResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type GatewayRuleNewResponseRuleSettingsBisoAdminControls struct {
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
	JSON gatewayRuleNewResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct [GatewayRuleNewResponseRuleSettingsBisoAdminControls]
type gatewayRuleNewResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type GatewayRuleNewResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                               `json:"enforce"`
	JSON    gatewayRuleNewResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsCheckSessionJSON contains the JSON metadata
// for the struct [GatewayRuleNewResponseRuleSettingsCheckSession]
type gatewayRuleNewResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleNewResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleNewResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleNewResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleNewResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsDNSResolversJSON contains the JSON metadata
// for the struct [GatewayRuleNewResponseRuleSettingsDNSResolvers]
type gatewayRuleNewResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleNewResponseRuleSettingsDNSResolversIPV4 struct {
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
	JSON   gatewayRuleNewResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsDNSResolversIPV4JSON contains the JSON
// metadata for the struct [GatewayRuleNewResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleNewResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsDNSResolversIPV4JSON) RawJSON() string {
	return r.raw
}

type GatewayRuleNewResponseRuleSettingsDNSResolversIPV6 struct {
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
	JSON   gatewayRuleNewResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsDNSResolversIPV6JSON contains the JSON
// metadata for the struct [GatewayRuleNewResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleNewResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsDNSResolversIPV6JSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleNewResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                       `json:"ipv6"`
	JSON gatewayRuleNewResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsEgressJSON contains the JSON metadata for the
// struct [GatewayRuleNewResponseRuleSettingsEgress]
type gatewayRuleNewResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleNewResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                            `json:"port"`
	JSON gatewayRuleNewResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsL4overrideJSON contains the JSON metadata for
// the struct [GatewayRuleNewResponseRuleSettingsL4override]
type gatewayRuleNewResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleNewResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                     `json:"support_url"`
	JSON       gatewayRuleNewResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsNotificationSettingsJSON contains the JSON
// metadata for the struct [GatewayRuleNewResponseRuleSettingsNotificationSettings]
type gatewayRuleNewResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type GatewayRuleNewResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                             `json:"enabled"`
	JSON    gatewayRuleNewResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsPayloadLogJSON contains the JSON metadata for
// the struct [GatewayRuleNewResponseRuleSettingsPayloadLog]
type gatewayRuleNewResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleNewResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleNewResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleNewResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleNewResponseRuleSettingsUntrustedCertJSON contains the JSON metadata
// for the struct [GatewayRuleNewResponseRuleSettingsUntrustedCert]
type gatewayRuleNewResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseRuleSettingsUntrustedCertJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleNewResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleNewResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleNewResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleNewResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleNewResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleNewResponseRuleSettingsUntrustedCertActionError       GatewayRuleNewResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleNewResponseSchedule struct {
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
	JSON gatewayRuleNewResponseScheduleJSON `json:"-"`
}

// gatewayRuleNewResponseScheduleJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseSchedule]
type gatewayRuleNewResponseScheduleJSON struct {
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

func (r *GatewayRuleNewResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseScheduleJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV4JSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsDNSResolversIPV6JSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsEgressJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseRuleSettingsUntrustedCertJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleListResponse struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    GatewayRuleListResponseAction `json:"action"`
	CreatedAt time.Time                     `json:"created_at" format:"date-time"`
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
	Filters []GatewayRuleListResponseFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings GatewayRuleListResponseRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule GatewayRuleListResponseSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                      `json:"traffic"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      gatewayRuleListResponseJSON `json:"-"`
}

// gatewayRuleListResponseJSON contains the JSON metadata for the struct
// [GatewayRuleListResponse]
type gatewayRuleListResponseJSON struct {
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

func (r *GatewayRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleListResponseAction string

const (
	GatewayRuleListResponseActionOn           GatewayRuleListResponseAction = "on"
	GatewayRuleListResponseActionOff          GatewayRuleListResponseAction = "off"
	GatewayRuleListResponseActionAllow        GatewayRuleListResponseAction = "allow"
	GatewayRuleListResponseActionBlock        GatewayRuleListResponseAction = "block"
	GatewayRuleListResponseActionScan         GatewayRuleListResponseAction = "scan"
	GatewayRuleListResponseActionNoscan       GatewayRuleListResponseAction = "noscan"
	GatewayRuleListResponseActionSafesearch   GatewayRuleListResponseAction = "safesearch"
	GatewayRuleListResponseActionYtrestricted GatewayRuleListResponseAction = "ytrestricted"
	GatewayRuleListResponseActionIsolate      GatewayRuleListResponseAction = "isolate"
	GatewayRuleListResponseActionNoisolate    GatewayRuleListResponseAction = "noisolate"
	GatewayRuleListResponseActionOverride     GatewayRuleListResponseAction = "override"
	GatewayRuleListResponseActionL4Override   GatewayRuleListResponseAction = "l4_override"
	GatewayRuleListResponseActionEgress       GatewayRuleListResponseAction = "egress"
	GatewayRuleListResponseActionAuditSSH     GatewayRuleListResponseAction = "audit_ssh"
)

// The protocol or layer to use.
type GatewayRuleListResponseFilter string

const (
	GatewayRuleListResponseFilterHTTP   GatewayRuleListResponseFilter = "http"
	GatewayRuleListResponseFilterDNS    GatewayRuleListResponseFilter = "dns"
	GatewayRuleListResponseFilterL4     GatewayRuleListResponseFilter = "l4"
	GatewayRuleListResponseFilterEgress GatewayRuleListResponseFilter = "egress"
)

// Additional settings that modify the rule's action.
type GatewayRuleListResponseRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH GatewayRuleListResponseRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls GatewayRuleListResponseRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession GatewayRuleListResponseRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers GatewayRuleListResponseRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress GatewayRuleListResponseRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override GatewayRuleListResponseRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings GatewayRuleListResponseRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog GatewayRuleListResponseRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert GatewayRuleListResponseRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          gatewayRuleListResponseRuleSettingsJSON          `json:"-"`
}

// gatewayRuleListResponseRuleSettingsJSON contains the JSON metadata for the
// struct [GatewayRuleListResponseRuleSettings]
type gatewayRuleListResponseRuleSettingsJSON struct {
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

func (r *GatewayRuleListResponseRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type GatewayRuleListResponseRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                            `json:"command_logging"`
	JSON           gatewayRuleListResponseRuleSettingsAuditSSHJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsAuditSSHJSON contains the JSON metadata for
// the struct [GatewayRuleListResponseRuleSettingsAuditSSH]
type gatewayRuleListResponseRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type GatewayRuleListResponseRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                     `json:"du"`
	JSON gatewayRuleListResponseRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct [GatewayRuleListResponseRuleSettingsBisoAdminControls]
type gatewayRuleListResponseRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type GatewayRuleListResponseRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                `json:"enforce"`
	JSON    gatewayRuleListResponseRuleSettingsCheckSessionJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsCheckSessionJSON contains the JSON metadata
// for the struct [GatewayRuleListResponseRuleSettingsCheckSession]
type gatewayRuleListResponseRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type GatewayRuleListResponseRuleSettingsDNSResolvers struct {
	IPV4 []GatewayRuleListResponseRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []GatewayRuleListResponseRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON gatewayRuleListResponseRuleSettingsDNSResolversJSON   `json:"-"`
}

// gatewayRuleListResponseRuleSettingsDNSResolversJSON contains the JSON metadata
// for the struct [GatewayRuleListResponseRuleSettingsDNSResolvers]
type gatewayRuleListResponseRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleListResponseRuleSettingsDNSResolversIPV4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                  `json:"vnet_id"`
	JSON   gatewayRuleListResponseRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsDNSResolversIPV4JSON contains the JSON
// metadata for the struct [GatewayRuleListResponseRuleSettingsDNSResolversIPV4]
type gatewayRuleListResponseRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsDNSResolversIPV4JSON) RawJSON() string {
	return r.raw
}

type GatewayRuleListResponseRuleSettingsDNSResolversIPV6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                  `json:"vnet_id"`
	JSON   gatewayRuleListResponseRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsDNSResolversIPV6JSON contains the JSON
// metadata for the struct [GatewayRuleListResponseRuleSettingsDNSResolversIPV6]
type gatewayRuleListResponseRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsDNSResolversIPV6JSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleListResponseRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                        `json:"ipv6"`
	JSON gatewayRuleListResponseRuleSettingsEgressJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsEgressJSON contains the JSON metadata for the
// struct [GatewayRuleListResponseRuleSettingsEgress]
type gatewayRuleListResponseRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleListResponseRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                             `json:"port"`
	JSON gatewayRuleListResponseRuleSettingsL4overrideJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsL4overrideJSON contains the JSON metadata for
// the struct [GatewayRuleListResponseRuleSettingsL4override]
type gatewayRuleListResponseRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleListResponseRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                      `json:"support_url"`
	JSON       gatewayRuleListResponseRuleSettingsNotificationSettingsJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsNotificationSettingsJSON contains the JSON
// metadata for the struct
// [GatewayRuleListResponseRuleSettingsNotificationSettings]
type gatewayRuleListResponseRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type GatewayRuleListResponseRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                              `json:"enabled"`
	JSON    gatewayRuleListResponseRuleSettingsPayloadLogJSON `json:"-"`
}

// gatewayRuleListResponseRuleSettingsPayloadLogJSON contains the JSON metadata for
// the struct [GatewayRuleListResponseRuleSettingsPayloadLog]
type gatewayRuleListResponseRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleListResponseRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action GatewayRuleListResponseRuleSettingsUntrustedCertAction `json:"action"`
	JSON   gatewayRuleListResponseRuleSettingsUntrustedCertJSON   `json:"-"`
}

// gatewayRuleListResponseRuleSettingsUntrustedCertJSON contains the JSON metadata
// for the struct [GatewayRuleListResponseRuleSettingsUntrustedCert]
type gatewayRuleListResponseRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseRuleSettingsUntrustedCertJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleListResponseRuleSettingsUntrustedCertAction string

const (
	GatewayRuleListResponseRuleSettingsUntrustedCertActionPassThrough GatewayRuleListResponseRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleListResponseRuleSettingsUntrustedCertActionBlock       GatewayRuleListResponseRuleSettingsUntrustedCertAction = "block"
	GatewayRuleListResponseRuleSettingsUntrustedCertActionError       GatewayRuleListResponseRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleListResponseSchedule struct {
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
	Wed  string                              `json:"wed"`
	JSON gatewayRuleListResponseScheduleJSON `json:"-"`
}

// gatewayRuleListResponseScheduleJSON contains the JSON metadata for the struct
// [GatewayRuleListResponseSchedule]
type gatewayRuleListResponseScheduleJSON struct {
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

func (r *GatewayRuleListResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseScheduleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.GatewayRuleDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayRuleDeleteResponse interface {
	ImplementsZeroTrustGatewayRuleDeleteResponse()
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

func (r gatewayRuleGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsDNSResolversIPV4JSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsDNSResolversIPV6JSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsEgressJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseRuleSettingsUntrustedCertJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
)

// The protocol or layer to use.
type GatewayRuleNewParamsFilter string

const (
	GatewayRuleNewParamsFilterHTTP   GatewayRuleNewParamsFilter = "http"
	GatewayRuleNewParamsFilterDNS    GatewayRuleNewParamsFilter = "dns"
	GatewayRuleNewParamsFilterL4     GatewayRuleNewParamsFilter = "l4"
	GatewayRuleNewParamsFilterEgress GatewayRuleNewParamsFilter = "egress"
)

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
	// route to the address closest to their origin.
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
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[GatewayRuleNewParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
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
// route to the address closest to their origin.
type GatewayRuleNewParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]GatewayRuleNewParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]GatewayRuleNewParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r GatewayRuleNewParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewParamsRuleSettingsDNSResolversIPV4 struct {
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

func (r GatewayRuleNewParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewParamsRuleSettingsDNSResolversIPV6 struct {
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
type GatewayRuleNewParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[GatewayRuleNewParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r GatewayRuleNewParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleNewParamsRuleSettingsUntrustedCertAction string

const (
	GatewayRuleNewParamsRuleSettingsUntrustedCertActionPassThrough GatewayRuleNewParamsRuleSettingsUntrustedCertAction = "pass_through"
	GatewayRuleNewParamsRuleSettingsUntrustedCertActionBlock       GatewayRuleNewParamsRuleSettingsUntrustedCertAction = "block"
	GatewayRuleNewParamsRuleSettingsUntrustedCertActionError       GatewayRuleNewParamsRuleSettingsUntrustedCertAction = "error"
)

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
	Errors   []GatewayRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayRuleNewResponse                   `json:"result,required"`
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

type GatewayRuleNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    gatewayRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayRuleNewResponseEnvelopeErrors]
type gatewayRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    gatewayRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayRuleNewResponseEnvelopeMessages]
type gatewayRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleNewResponseEnvelopeSuccess bool

const (
	GatewayRuleNewResponseEnvelopeSuccessTrue GatewayRuleNewResponseEnvelopeSuccess = true
)

type GatewayRuleUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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

func (r gatewayRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	GatewayRuleUpdateResponseEnvelopeSuccessTrue GatewayRuleUpdateResponseEnvelopeSuccess = true
)

type GatewayRuleListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayRuleListResponseEnvelope struct {
	Errors   []GatewayRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayRuleListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayRuleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayRuleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayRuleListResponseEnvelopeJSON       `json:"-"`
}

// gatewayRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleListResponseEnvelope]
type gatewayRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    gatewayRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayRuleListResponseEnvelopeErrors]
type gatewayRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayRuleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayRuleListResponseEnvelopeMessages]
type gatewayRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleListResponseEnvelopeSuccess bool

const (
	GatewayRuleListResponseEnvelopeSuccessTrue GatewayRuleListResponseEnvelopeSuccess = true
)

type GatewayRuleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       gatewayRuleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayRuleListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [GatewayRuleListResponseEnvelopeResultInfo]
type gatewayRuleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type GatewayRuleDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

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

func (r gatewayRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	GatewayRuleDeleteResponseEnvelopeSuccessTrue GatewayRuleDeleteResponseEnvelopeSuccess = true
)

type GatewayRuleGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

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

func (r gatewayRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleGetResponseEnvelopeSuccess bool

const (
	GatewayRuleGetResponseEnvelopeSuccessTrue GatewayRuleGetResponseEnvelopeSuccess = true
)

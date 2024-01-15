// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayRuleService] method
// instead.
type AccountGatewayRuleService struct {
	Options []option.RequestOption
}

// NewAccountGatewayRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayRuleService(opts ...option.RequestOption) (r *AccountGatewayRuleService) {
	r = &AccountGatewayRuleService{}
	r.Options = opts
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountGatewayRuleUpdateParams, opts ...option.RequestOption) (res *AccountGatewayRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) ZeroTrustGatewayRulesNewZeroTrustGatewayRule(ctx context.Context, identifier interface{}, body AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams, opts ...option.RequestOption) (res *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *AccountGatewayRuleService) ZeroTrustGatewayRulesListZeroTrustGatewayRules(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayRuleGetResponse struct {
	Errors   []AccountGatewayRuleGetResponseError   `json:"errors"`
	Messages []AccountGatewayRuleGetResponseMessage `json:"messages"`
	Result   AccountGatewayRuleGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayRuleGetResponseSuccess `json:"success"`
	JSON    accountGatewayRuleGetResponseJSON    `json:"-"`
}

// accountGatewayRuleGetResponseJSON contains the JSON metadata for the struct
// [AccountGatewayRuleGetResponse]
type accountGatewayRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountGatewayRuleGetResponseErrorJSON `json:"-"`
}

// accountGatewayRuleGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountGatewayRuleGetResponseError]
type accountGatewayRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountGatewayRuleGetResponseMessageJSON `json:"-"`
}

// accountGatewayRuleGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayRuleGetResponseMessage]
type accountGatewayRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleGetResponseResult struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    AccountGatewayRuleGetResponseResultAction `json:"action"`
	CreatedAt time.Time                                 `json:"created_at" format:"date-time"`
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
	Filters []AccountGatewayRuleGetResponseResultFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings AccountGatewayRuleGetResponseResultRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule AccountGatewayRuleGetResponseResultSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                  `json:"traffic"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accountGatewayRuleGetResponseResultJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayRuleGetResponseResult]
type accountGatewayRuleGetResponseResultJSON struct {
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

func (r *AccountGatewayRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleGetResponseResultAction string

const (
	AccountGatewayRuleGetResponseResultActionOn           AccountGatewayRuleGetResponseResultAction = "on"
	AccountGatewayRuleGetResponseResultActionOff          AccountGatewayRuleGetResponseResultAction = "off"
	AccountGatewayRuleGetResponseResultActionAllow        AccountGatewayRuleGetResponseResultAction = "allow"
	AccountGatewayRuleGetResponseResultActionBlock        AccountGatewayRuleGetResponseResultAction = "block"
	AccountGatewayRuleGetResponseResultActionScan         AccountGatewayRuleGetResponseResultAction = "scan"
	AccountGatewayRuleGetResponseResultActionNoscan       AccountGatewayRuleGetResponseResultAction = "noscan"
	AccountGatewayRuleGetResponseResultActionSafesearch   AccountGatewayRuleGetResponseResultAction = "safesearch"
	AccountGatewayRuleGetResponseResultActionYtrestricted AccountGatewayRuleGetResponseResultAction = "ytrestricted"
	AccountGatewayRuleGetResponseResultActionIsolate      AccountGatewayRuleGetResponseResultAction = "isolate"
	AccountGatewayRuleGetResponseResultActionNoisolate    AccountGatewayRuleGetResponseResultAction = "noisolate"
	AccountGatewayRuleGetResponseResultActionOverride     AccountGatewayRuleGetResponseResultAction = "override"
	AccountGatewayRuleGetResponseResultActionL4Override   AccountGatewayRuleGetResponseResultAction = "l4_override"
	AccountGatewayRuleGetResponseResultActionEgress       AccountGatewayRuleGetResponseResultAction = "egress"
	AccountGatewayRuleGetResponseResultActionAuditSSH     AccountGatewayRuleGetResponseResultAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleGetResponseResultFilter string

const (
	AccountGatewayRuleGetResponseResultFilterHTTP   AccountGatewayRuleGetResponseResultFilter = "http"
	AccountGatewayRuleGetResponseResultFilterDNS    AccountGatewayRuleGetResponseResultFilter = "dns"
	AccountGatewayRuleGetResponseResultFilterL4     AccountGatewayRuleGetResponseResultFilter = "l4"
	AccountGatewayRuleGetResponseResultFilterEgress AccountGatewayRuleGetResponseResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleGetResponseResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH AccountGatewayRuleGetResponseResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls AccountGatewayRuleGetResponseResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession AccountGatewayRuleGetResponseResultRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers AccountGatewayRuleGetResponseResultRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress AccountGatewayRuleGetResponseResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override AccountGatewayRuleGetResponseResultRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings AccountGatewayRuleGetResponseResultRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog AccountGatewayRuleGetResponseResultRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          accountGatewayRuleGetResponseResultRuleSettingsJSON          `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsJSON contains the JSON metadata
// for the struct [AccountGatewayRuleGetResponseResultRuleSettings]
type accountGatewayRuleGetResponseResultRuleSettingsJSON struct {
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
	InsecureDisableDnssecValidation apijson.Field
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

func (r *AccountGatewayRuleGetResponseResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleGetResponseResultRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                        `json:"command_logging"`
	JSON           accountGatewayRuleGetResponseResultRuleSettingsAuditSSHJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsAuditSSHJSON contains the JSON
// metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsAuditSSH]
type accountGatewayRuleGetResponseResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleGetResponseResultRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                 `json:"du"`
	JSON accountGatewayRuleGetResponseResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsBisoAdminControlsJSON contains
// the JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsBisoAdminControls]
type accountGatewayRuleGetResponseResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type AccountGatewayRuleGetResponseResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                            `json:"enforce"`
	JSON    accountGatewayRuleGetResponseResultRuleSettingsCheckSessionJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsCheckSessionJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsCheckSession]
type accountGatewayRuleGetResponseResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleGetResponseResultRuleSettingsDNSResolvers struct {
	Ipv4 []AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4 `json:"ipv4"`
	Ipv6 []AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6 `json:"ipv6"`
	JSON accountGatewayRuleGetResponseResultRuleSettingsDNSResolversJSON   `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsDNSResolversJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsDNSResolvers]
type accountGatewayRuleGetResponseResultRuleSettingsDNSResolversJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                              `json:"vnet_id"`
	JSON   accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4JSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4JSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4]
type accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                              `json:"vnet_id"`
	JSON   accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6JSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6JSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6]
type accountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsDNSResolversIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleGetResponseResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                                    `json:"ipv6"`
	JSON accountGatewayRuleGetResponseResultRuleSettingsEgressJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsEgressJSON contains the JSON
// metadata for the struct [AccountGatewayRuleGetResponseResultRuleSettingsEgress]
type accountGatewayRuleGetResponseResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleGetResponseResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                         `json:"port"`
	JSON accountGatewayRuleGetResponseResultRuleSettingsL4overrideJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsL4overrideJSON contains the JSON
// metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsL4override]
type accountGatewayRuleGetResponseResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleGetResponseResultRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                  `json:"support_url"`
	JSON       accountGatewayRuleGetResponseResultRuleSettingsNotificationSettingsJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsNotificationSettingsJSON contains
// the JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsNotificationSettings]
type accountGatewayRuleGetResponseResultRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type AccountGatewayRuleGetResponseResultRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                          `json:"enabled"`
	JSON    accountGatewayRuleGetResponseResultRuleSettingsPayloadLogJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsPayloadLogJSON contains the JSON
// metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsPayloadLog]
type accountGatewayRuleGetResponseResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   accountGatewayRuleGetResponseResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// accountGatewayRuleGetResponseResultRuleSettingsUntrustedCertJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCert]
type accountGatewayRuleGetResponseResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertActionError       AccountGatewayRuleGetResponseResultRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleGetResponseResultSchedule struct {
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
	Wed  string                                          `json:"wed"`
	JSON accountGatewayRuleGetResponseResultScheduleJSON `json:"-"`
}

// accountGatewayRuleGetResponseResultScheduleJSON contains the JSON metadata for
// the struct [AccountGatewayRuleGetResponseResultSchedule]
type accountGatewayRuleGetResponseResultScheduleJSON struct {
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

func (r *AccountGatewayRuleGetResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayRuleGetResponseSuccess bool

const (
	AccountGatewayRuleGetResponseSuccessTrue AccountGatewayRuleGetResponseSuccess = true
)

type AccountGatewayRuleUpdateResponse struct {
	Errors   []AccountGatewayRuleUpdateResponseError   `json:"errors"`
	Messages []AccountGatewayRuleUpdateResponseMessage `json:"messages"`
	Result   AccountGatewayRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayRuleUpdateResponseSuccess `json:"success"`
	JSON    accountGatewayRuleUpdateResponseJSON    `json:"-"`
}

// accountGatewayRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AccountGatewayRuleUpdateResponse]
type accountGatewayRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountGatewayRuleUpdateResponseErrorJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayRuleUpdateResponseError]
type accountGatewayRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountGatewayRuleUpdateResponseMessageJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayRuleUpdateResponseMessage]
type accountGatewayRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleUpdateResponseResult struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    AccountGatewayRuleUpdateResponseResultAction `json:"action"`
	CreatedAt time.Time                                    `json:"created_at" format:"date-time"`
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
	Filters []AccountGatewayRuleUpdateResponseResultFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings AccountGatewayRuleUpdateResponseResultRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule AccountGatewayRuleUpdateResponseResultSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                     `json:"traffic"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountGatewayRuleUpdateResponseResultJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayRuleUpdateResponseResult]
type accountGatewayRuleUpdateResponseResultJSON struct {
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

func (r *AccountGatewayRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleUpdateResponseResultAction string

const (
	AccountGatewayRuleUpdateResponseResultActionOn           AccountGatewayRuleUpdateResponseResultAction = "on"
	AccountGatewayRuleUpdateResponseResultActionOff          AccountGatewayRuleUpdateResponseResultAction = "off"
	AccountGatewayRuleUpdateResponseResultActionAllow        AccountGatewayRuleUpdateResponseResultAction = "allow"
	AccountGatewayRuleUpdateResponseResultActionBlock        AccountGatewayRuleUpdateResponseResultAction = "block"
	AccountGatewayRuleUpdateResponseResultActionScan         AccountGatewayRuleUpdateResponseResultAction = "scan"
	AccountGatewayRuleUpdateResponseResultActionNoscan       AccountGatewayRuleUpdateResponseResultAction = "noscan"
	AccountGatewayRuleUpdateResponseResultActionSafesearch   AccountGatewayRuleUpdateResponseResultAction = "safesearch"
	AccountGatewayRuleUpdateResponseResultActionYtrestricted AccountGatewayRuleUpdateResponseResultAction = "ytrestricted"
	AccountGatewayRuleUpdateResponseResultActionIsolate      AccountGatewayRuleUpdateResponseResultAction = "isolate"
	AccountGatewayRuleUpdateResponseResultActionNoisolate    AccountGatewayRuleUpdateResponseResultAction = "noisolate"
	AccountGatewayRuleUpdateResponseResultActionOverride     AccountGatewayRuleUpdateResponseResultAction = "override"
	AccountGatewayRuleUpdateResponseResultActionL4Override   AccountGatewayRuleUpdateResponseResultAction = "l4_override"
	AccountGatewayRuleUpdateResponseResultActionEgress       AccountGatewayRuleUpdateResponseResultAction = "egress"
	AccountGatewayRuleUpdateResponseResultActionAuditSSH     AccountGatewayRuleUpdateResponseResultAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleUpdateResponseResultFilter string

const (
	AccountGatewayRuleUpdateResponseResultFilterHTTP   AccountGatewayRuleUpdateResponseResultFilter = "http"
	AccountGatewayRuleUpdateResponseResultFilterDNS    AccountGatewayRuleUpdateResponseResultFilter = "dns"
	AccountGatewayRuleUpdateResponseResultFilterL4     AccountGatewayRuleUpdateResponseResultFilter = "l4"
	AccountGatewayRuleUpdateResponseResultFilterEgress AccountGatewayRuleUpdateResponseResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleUpdateResponseResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH AccountGatewayRuleUpdateResponseResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls AccountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession AccountGatewayRuleUpdateResponseResultRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress AccountGatewayRuleUpdateResponseResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override AccountGatewayRuleUpdateResponseResultRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings AccountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog AccountGatewayRuleUpdateResponseResultRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          accountGatewayRuleUpdateResponseResultRuleSettingsJSON          `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsJSON contains the JSON
// metadata for the struct [AccountGatewayRuleUpdateResponseResultRuleSettings]
type accountGatewayRuleUpdateResponseResultRuleSettingsJSON struct {
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
	InsecureDisableDnssecValidation apijson.Field
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

func (r *AccountGatewayRuleUpdateResponseResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleUpdateResponseResultRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                           `json:"command_logging"`
	JSON           accountGatewayRuleUpdateResponseResultRuleSettingsAuditSSHJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsAuditSSHJSON contains the JSON
// metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsAuditSSH]
type accountGatewayRuleUpdateResponseResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                    `json:"du"`
	JSON accountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControlsJSON contains
// the JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControls]
type accountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type AccountGatewayRuleUpdateResponseResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                               `json:"enforce"`
	JSON    accountGatewayRuleUpdateResponseResultRuleSettingsCheckSessionJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsCheckSessionJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsCheckSession]
type accountGatewayRuleUpdateResponseResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolvers struct {
	Ipv4 []AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4 `json:"ipv4"`
	Ipv6 []AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6 `json:"ipv6"`
	JSON accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversJSON   `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolvers]
type accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                 `json:"vnet_id"`
	JSON   accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4JSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4JSON contains
// the JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4]
type accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                 `json:"vnet_id"`
	JSON   accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6JSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6JSON contains
// the JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6]
type accountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsDNSResolversIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleUpdateResponseResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                                       `json:"ipv6"`
	JSON accountGatewayRuleUpdateResponseResultRuleSettingsEgressJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsEgressJSON contains the JSON
// metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsEgress]
type accountGatewayRuleUpdateResponseResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleUpdateResponseResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                            `json:"port"`
	JSON accountGatewayRuleUpdateResponseResultRuleSettingsL4overrideJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsL4overrideJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsL4override]
type accountGatewayRuleUpdateResponseResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                     `json:"support_url"`
	JSON       accountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettingsJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettings]
type accountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type AccountGatewayRuleUpdateResponseResultRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                             `json:"enabled"`
	JSON    accountGatewayRuleUpdateResponseResultRuleSettingsPayloadLogJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsPayloadLogJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsPayloadLog]
type accountGatewayRuleUpdateResponseResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   accountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// accountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertJSON contains the
// JSON metadata for the struct
// [AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCert]
type accountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertActionError       AccountGatewayRuleUpdateResponseResultRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleUpdateResponseResultSchedule struct {
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
	Wed  string                                             `json:"wed"`
	JSON accountGatewayRuleUpdateResponseResultScheduleJSON `json:"-"`
}

// accountGatewayRuleUpdateResponseResultScheduleJSON contains the JSON metadata
// for the struct [AccountGatewayRuleUpdateResponseResultSchedule]
type accountGatewayRuleUpdateResponseResultScheduleJSON struct {
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

func (r *AccountGatewayRuleUpdateResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayRuleUpdateResponseSuccess bool

const (
	AccountGatewayRuleUpdateResponseSuccessTrue AccountGatewayRuleUpdateResponseSuccess = true
)

type AccountGatewayRuleDeleteResponse struct {
	Errors   []AccountGatewayRuleDeleteResponseError   `json:"errors"`
	Messages []AccountGatewayRuleDeleteResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayRuleDeleteResponseSuccess `json:"success"`
	JSON    accountGatewayRuleDeleteResponseJSON    `json:"-"`
}

// accountGatewayRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountGatewayRuleDeleteResponse]
type accountGatewayRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountGatewayRuleDeleteResponseErrorJSON `json:"-"`
}

// accountGatewayRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayRuleDeleteResponseError]
type accountGatewayRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountGatewayRuleDeleteResponseMessageJSON `json:"-"`
}

// accountGatewayRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayRuleDeleteResponseMessage]
type accountGatewayRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayRuleDeleteResponseSuccess bool

const (
	AccountGatewayRuleDeleteResponseSuccessTrue AccountGatewayRuleDeleteResponseSuccess = true
)

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse struct {
	Errors   []AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseError   `json:"errors"`
	Messages []AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessage `json:"messages"`
	Result   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSuccess `json:"success"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON    `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseErrorJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseError]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessageJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessage]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResult struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction `json:"action"`
	CreatedAt time.Time                                                                          `json:"created_at" format:"date-time"`
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
	Filters []AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                                                           `json:"traffic"`
	UpdatedAt time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON      accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResult]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultJSON struct {
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionOn           AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "on"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionOff          AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "off"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionAllow        AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "allow"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionBlock        AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionScan         AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "scan"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionNoscan       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "noscan"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionSafesearch   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "safesearch"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionYtrestricted AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "ytrestricted"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionIsolate      AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "isolate"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionNoisolate    AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "noisolate"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionOverride     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "override"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionL4Override   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "l4_override"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionEgress       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "egress"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultActionAuditSSH     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilterHTTP   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter = "http"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilterDNS    AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter = "dns"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilterL4     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter = "l4"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilterEgress AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsJSON          `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettings]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsJSON struct {
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
	InsecureDisableDnssecValidation apijson.Field
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                                                                 `json:"command_logging"`
	JSON           accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSHJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSHJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSH]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                                                          `json:"du"`
	JSON accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControlsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControls]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                                                                     `json:"enforce"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSessionJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSessionJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSession]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolvers struct {
	Ipv4 []AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4 `json:"ipv4"`
	Ipv6 []AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6 `json:"ipv6"`
	JSON accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversJSON   `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolvers]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                                       `json:"vnet_id"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4JSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4JSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                                       `json:"vnet_id"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6JSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6JSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsDNSResolversIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                                                                             `json:"ipv6"`
	JSON accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgressJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgressJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgress]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                                                                  `json:"port"`
	JSON accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4overrideJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4overrideJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4override]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                           `json:"support_url"`
	JSON       accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettingsJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettings]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                                                                   `json:"enabled"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLogJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLogJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLog]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCert]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertActionError       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultSchedule struct {
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
	Wed  string                                                                                   `json:"wed"`
	JSON accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultScheduleJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultScheduleJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultSchedule]
type accountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultScheduleJSON struct {
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSuccess bool

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSuccessTrue AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleResponseSuccess = true
)

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse struct {
	Errors     []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseError    `json:"errors"`
	Messages   []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessage  `json:"messages"`
	Result     []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResult   `json:"result"`
	ResultInfo AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSuccess `json:"success"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON    `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseErrorJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseError]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessageJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessage]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResult struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction `json:"action"`
	CreatedAt time.Time                                                                            `json:"created_at" format:"date-time"`
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
	Filters []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                                                                             `json:"traffic"`
	UpdatedAt time.Time                                                                          `json:"updated_at" format:"date-time"`
	JSON      accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResult]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultJSON struct {
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionOn           AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "on"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionOff          AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "off"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionAllow        AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "allow"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionBlock        AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionScan         AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "scan"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionNoscan       AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "noscan"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionSafesearch   AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "safesearch"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionYtrestricted AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "ytrestricted"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionIsolate      AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "isolate"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionNoisolate    AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "noisolate"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionOverride     AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "override"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionL4Override   AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "l4_override"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionEgress       AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "egress"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultActionAuditSSH     AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter string

const (
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilterHTTP   AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter = "http"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilterDNS    AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter = "dns"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilterL4     AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter = "l4"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilterEgress AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsJSON          `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettings]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsJSON struct {
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
	InsecureDisableDnssecValidation apijson.Field
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                                                                                   `json:"command_logging"`
	JSON           accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSHJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSHJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSH]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                                                                            `json:"du"`
	JSON accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControlsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControls]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                                                                                       `json:"enforce"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSessionJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSessionJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSession]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolvers struct {
	Ipv4 []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4 `json:"ipv4"`
	Ipv6 []AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6 `json:"ipv6"`
	JSON accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversJSON   `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolvers]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                                         `json:"vnet_id"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4JSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4JSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6 struct {
	// IP address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                                                                         `json:"vnet_id"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6JSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6JSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsDNSResolversIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                                                                               `json:"ipv6"`
	JSON accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgressJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgressJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgress]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                                                                    `json:"port"`
	JSON accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4overrideJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4overrideJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4override]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                             `json:"support_url"`
	JSON       accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettingsJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettings]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                                                                                     `json:"enabled"`
	JSON    accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLogJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLogJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLog]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCert]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertActionError       AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultSchedule struct {
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
	Wed  string                                                                                     `json:"wed"`
	JSON accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultScheduleJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultScheduleJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultSchedule]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultScheduleJSON struct {
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

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                `json:"total_count"`
	JSON       accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfoJSON `json:"-"`
}

// accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfo]
type accountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSuccess bool

const (
	AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSuccessTrue AccountGatewayRuleZeroTrustGatewayRulesListZeroTrustGatewayRulesResponseSuccess = true
)

type AccountGatewayRuleUpdateParams struct {
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[AccountGatewayRuleUpdateParamsAction] `json:"action,required"`
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
	Filters param.Field[[]AccountGatewayRuleUpdateParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[AccountGatewayRuleUpdateParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[AccountGatewayRuleUpdateParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleUpdateParamsAction string

const (
	AccountGatewayRuleUpdateParamsActionOn           AccountGatewayRuleUpdateParamsAction = "on"
	AccountGatewayRuleUpdateParamsActionOff          AccountGatewayRuleUpdateParamsAction = "off"
	AccountGatewayRuleUpdateParamsActionAllow        AccountGatewayRuleUpdateParamsAction = "allow"
	AccountGatewayRuleUpdateParamsActionBlock        AccountGatewayRuleUpdateParamsAction = "block"
	AccountGatewayRuleUpdateParamsActionScan         AccountGatewayRuleUpdateParamsAction = "scan"
	AccountGatewayRuleUpdateParamsActionNoscan       AccountGatewayRuleUpdateParamsAction = "noscan"
	AccountGatewayRuleUpdateParamsActionSafesearch   AccountGatewayRuleUpdateParamsAction = "safesearch"
	AccountGatewayRuleUpdateParamsActionYtrestricted AccountGatewayRuleUpdateParamsAction = "ytrestricted"
	AccountGatewayRuleUpdateParamsActionIsolate      AccountGatewayRuleUpdateParamsAction = "isolate"
	AccountGatewayRuleUpdateParamsActionNoisolate    AccountGatewayRuleUpdateParamsAction = "noisolate"
	AccountGatewayRuleUpdateParamsActionOverride     AccountGatewayRuleUpdateParamsAction = "override"
	AccountGatewayRuleUpdateParamsActionL4Override   AccountGatewayRuleUpdateParamsAction = "l4_override"
	AccountGatewayRuleUpdateParamsActionEgress       AccountGatewayRuleUpdateParamsAction = "egress"
	AccountGatewayRuleUpdateParamsActionAuditSSH     AccountGatewayRuleUpdateParamsAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleUpdateParamsFilter string

const (
	AccountGatewayRuleUpdateParamsFilterHTTP   AccountGatewayRuleUpdateParamsFilter = "http"
	AccountGatewayRuleUpdateParamsFilterDNS    AccountGatewayRuleUpdateParamsFilter = "dns"
	AccountGatewayRuleUpdateParamsFilterL4     AccountGatewayRuleUpdateParamsFilter = "l4"
	AccountGatewayRuleUpdateParamsFilterEgress AccountGatewayRuleUpdateParamsFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleUpdateParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[AccountGatewayRuleUpdateParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers param.Field[AccountGatewayRuleUpdateParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[AccountGatewayRuleUpdateParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[AccountGatewayRuleUpdateParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[AccountGatewayRuleUpdateParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls struct {
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

func (r AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type AccountGatewayRuleUpdateParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleUpdateParamsRuleSettingsDNSResolvers struct {
	Ipv4 param.Field[[]AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv4] `json:"ipv4"`
	Ipv6 param.Field[[]AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv6] `json:"ipv6"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv4 struct {
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

func (r AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv6 struct {
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

func (r AccountGatewayRuleUpdateParamsRuleSettingsDNSResolversIpv6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleUpdateParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 param.Field[string] `json:"ipv6"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleUpdateParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleUpdateParamsRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError       AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleUpdateParamsSchedule struct {
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

func (r AccountGatewayRuleUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams struct {
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction] `json:"action,required"`
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
	Filters param.Field[[]AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOn           AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "on"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOff          AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "off"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionAllow        AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "allow"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionBlock        AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionScan         AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "scan"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionNoscan       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "noscan"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionSafesearch   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "safesearch"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionYtrestricted AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "ytrestricted"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionIsolate      AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "isolate"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionNoisolate    AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "noisolate"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionOverride     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "override"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionL4Override   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "l4_override"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionEgress       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "egress"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsActionAuditSSH     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction = "audit_ssh"
)

// The protocol or layer to use.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterHTTP   AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "http"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterDNS    AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "dns"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterL4     AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "l4"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilterEgress AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter = "egress"
)

// Additional settings that modify the rule's action.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin.
	DNSResolvers param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls struct {
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

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers struct {
	Ipv4 param.Field[[]AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv4] `json:"ipv4"`
	Ipv6 param.Field[[]AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv6] `json:"ipv6"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv4 struct {
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

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv6 struct {
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

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsDNSResolversIpv6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 param.Field[string] `json:"ipv6"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings struct {
	// set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionError       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "error"
)

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule struct {
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

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

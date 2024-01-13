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

// Fetch a single Zero Trust Gateway Rule.
func (r *AccountGatewayRuleService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SingleResponseH8mJb2Ar, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured Zero Trust Gateway Rule.
func (r *AccountGatewayRuleService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountGatewayRuleUpdateParams, opts ...option.RequestOption) (res *SingleResponseH8mJb2Ar, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Zero Trust Gateway Rule.
func (r *AccountGatewayRuleService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *EmptyResponseArJnNcMr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Zero Trust Gateway Rule.
func (r *AccountGatewayRuleService) ZeroTrustGatewayRulesNewZeroTrustGatewayRule(ctx context.Context, identifier interface{}, body AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams, opts ...option.RequestOption) (res *SingleResponseH8mJb2Ar, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Zero Trust Gateway Rules for an account.
func (r *AccountGatewayRuleService) ZeroTrustGatewayRulesListZeroTrustGatewayRules(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionVRHfxGdB, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/rules", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionVRHfxGdB struct {
	Errors     []ResponseCollectionVRHfxGdBError    `json:"errors"`
	Messages   []ResponseCollectionVRHfxGdBMessage  `json:"messages"`
	Result     []ResponseCollectionVRHfxGdBResult   `json:"result"`
	ResultInfo ResponseCollectionVRHfxGdBResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionVRHfxGdBSuccess `json:"success"`
	JSON    responseCollectionVrHfxGdBJSON    `json:"-"`
}

// responseCollectionVrHfxGdBJSON contains the JSON metadata for the struct
// [ResponseCollectionVRHfxGdB]
type responseCollectionVrHfxGdBJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVRHfxGdBError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionVrHfxGdBErrorJSON `json:"-"`
}

// responseCollectionVrHfxGdBErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionVRHfxGdBError]
type responseCollectionVrHfxGdBErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVRHfxGdBMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionVrHfxGdBMessageJSON `json:"-"`
}

// responseCollectionVrHfxGdBMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionVRHfxGdBMessage]
type responseCollectionVrHfxGdBMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVRHfxGdBResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to 'true'.
	Action    ResponseCollectionVRHfxGdBResultAction `json:"action"`
	CreatedAt time.Time                              `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the Rule.
	Description string `json:"description"`
	// The wirefilter expression to be used for device posture check matching.
	DevicePosture string `json:"device_posture"`
	// Set if the rule is enabled.
	Enabled bool `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []ResponseCollectionVRHfxGdBResultFilter `json:"filters"`
	// The wirefilter expression to be used for identity matching.
	Identity string `json:"identity"`
	// The name of the Rule.
	Name string `json:"name"`
	// Precedence sets the ordering of the rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings ResponseCollectionVRHfxGdBResultRuleSettings `json:"rule_settings"`
	// Schedule for activating DNS policies. Does not apply to HTTP or network
	// policies.
	Schedule ResponseCollectionVRHfxGdBResultSchedule `json:"schedule"`
	// The wirefilter expression to be used for traffic matching.
	Traffic   string                               `json:"traffic"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      responseCollectionVrHfxGdBResultJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultJSON contains the JSON metadata for the struct
// [ResponseCollectionVRHfxGdBResult]
type responseCollectionVrHfxGdBResultJSON struct {
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

func (r *ResponseCollectionVRHfxGdBResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to 'true'.
type ResponseCollectionVRHfxGdBResultAction string

const (
	ResponseCollectionVRHfxGdBResultActionOn           ResponseCollectionVRHfxGdBResultAction = "on"
	ResponseCollectionVRHfxGdBResultActionOff          ResponseCollectionVRHfxGdBResultAction = "off"
	ResponseCollectionVRHfxGdBResultActionAllow        ResponseCollectionVRHfxGdBResultAction = "allow"
	ResponseCollectionVRHfxGdBResultActionBlock        ResponseCollectionVRHfxGdBResultAction = "block"
	ResponseCollectionVRHfxGdBResultActionScan         ResponseCollectionVRHfxGdBResultAction = "scan"
	ResponseCollectionVRHfxGdBResultActionNoscan       ResponseCollectionVRHfxGdBResultAction = "noscan"
	ResponseCollectionVRHfxGdBResultActionSafesearch   ResponseCollectionVRHfxGdBResultAction = "safesearch"
	ResponseCollectionVRHfxGdBResultActionYtrestricted ResponseCollectionVRHfxGdBResultAction = "ytrestricted"
	ResponseCollectionVRHfxGdBResultActionIsolate      ResponseCollectionVRHfxGdBResultAction = "isolate"
	ResponseCollectionVRHfxGdBResultActionNoisolate    ResponseCollectionVRHfxGdBResultAction = "noisolate"
	ResponseCollectionVRHfxGdBResultActionOverride     ResponseCollectionVRHfxGdBResultAction = "override"
	ResponseCollectionVRHfxGdBResultActionL4Override   ResponseCollectionVRHfxGdBResultAction = "l4_override"
	ResponseCollectionVRHfxGdBResultActionEgress       ResponseCollectionVRHfxGdBResultAction = "egress"
	ResponseCollectionVRHfxGdBResultActionAuditSSH     ResponseCollectionVRHfxGdBResultAction = "audit_ssh"
)

// The protocol or layer to use.
type ResponseCollectionVRHfxGdBResultFilter string

const (
	ResponseCollectionVRHfxGdBResultFilterHTTP   ResponseCollectionVRHfxGdBResultFilter = "http"
	ResponseCollectionVRHfxGdBResultFilterDNS    ResponseCollectionVRHfxGdBResultFilter = "dns"
	ResponseCollectionVRHfxGdBResultFilterL4     ResponseCollectionVRHfxGdBResultFilter = "l4"
	ResponseCollectionVRHfxGdBResultFilterEgress ResponseCollectionVRHfxGdBResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type ResponseCollectionVRHfxGdBResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Audit ssh action settings
	AuditSSH ResponseCollectionVRHfxGdBResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls ResponseCollectionVRHfxGdBResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred that will be displayed on the custom
	// block page (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession ResponseCollectionVRHfxGdBResultRuleSettingsCheckSession `json:"check_session"`
	// Configure how Proxy traffic egresses. Can be set for rules with Egress action
	// and Egress filter. Can be omitted to indicate local egress via Warp IPs
	Egress ResponseCollectionVRHfxGdBResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Include IPs in DNS resolver category blocks. By default categories only block on
	// domain names.
	IPCategories bool `json:"ip_categories"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override ResponseCollectionVRHfxGdBResultRuleSettingsL4override `json:"l4override"`
	// Override matching DNS queries with this.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with this.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog ResponseCollectionVRHfxGdBResultRuleSettingsPayloadLog `json:"payload_log"`
	// Configure behavior when an upstream cert is invalid / an SSL error occurs.
	UntrustedCert ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          responseCollectionVrHfxGdBResultRuleSettingsJSON          `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsJSON contains the JSON metadata for
// the struct [ResponseCollectionVRHfxGdBResultRuleSettings]
type responseCollectionVrHfxGdBResultRuleSettingsJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BisoAdminControls               apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	Egress                          apijson.Field
	InsecureDisableDnssecValidation apijson.Field
	IPCategories                    apijson.Field
	L4override                      apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	UntrustedCert                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Audit ssh action settings
type ResponseCollectionVRHfxGdBResultRuleSettingsAuditSSH struct {
	// Turn on SSH command logging.
	CommandLogging bool                                                     `json:"command_logging"`
	JSON           responseCollectionVrHfxGdBResultRuleSettingsAuditSSHJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsAuditSSHJSON contains the JSON
// metadata for the struct [ResponseCollectionVRHfxGdBResultRuleSettingsAuditSSH]
type responseCollectionVrHfxGdBResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type ResponseCollectionVRHfxGdBResultRuleSettingsBisoAdminControls struct {
	// Disable copy-paste.
	Dcp bool `json:"dcp"`
	// Disable download.
	Dd bool `json:"dd"`
	// Disable keyboard usage.
	Dk bool `json:"dk"`
	// Disable printing.
	Dp bool `json:"dp"`
	// Disable upload.
	Du   bool                                                              `json:"du"`
	JSON responseCollectionVrHfxGdBResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsBisoAdminControlsJSON contains the
// JSON metadata for the struct
// [ResponseCollectionVRHfxGdBResultRuleSettingsBisoAdminControls]
type responseCollectionVrHfxGdBResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type ResponseCollectionVRHfxGdBResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Enable session enforcement for this fule.
	Enforce bool                                                         `json:"enforce"`
	JSON    responseCollectionVrHfxGdBResultRuleSettingsCheckSessionJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsCheckSessionJSON contains the JSON
// metadata for the struct
// [ResponseCollectionVRHfxGdBResultRuleSettingsCheckSession]
type responseCollectionVrHfxGdBResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Proxy traffic egresses. Can be set for rules with Egress action
// and Egress filter. Can be omitted to indicate local egress via Warp IPs
type ResponseCollectionVRHfxGdBResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The IPv4 address to be used for egress in the event of an error egressing with
	// the primary IPv4. Can be '0.0.0.0' to indicate local egreass via Warp IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                                 `json:"ipv6"`
	JSON responseCollectionVrHfxGdBResultRuleSettingsEgressJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsEgressJSON contains the JSON
// metadata for the struct [ResponseCollectionVRHfxGdBResultRuleSettingsEgress]
type responseCollectionVrHfxGdBResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type ResponseCollectionVRHfxGdBResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                      `json:"port"`
	JSON responseCollectionVrHfxGdBResultRuleSettingsL4overrideJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsL4overrideJSON contains the JSON
// metadata for the struct [ResponseCollectionVRHfxGdBResultRuleSettingsL4override]
type responseCollectionVrHfxGdBResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type ResponseCollectionVRHfxGdBResultRuleSettingsPayloadLog struct {
	// Enable DLP payload logging for this rule.
	Enabled bool                                                       `json:"enabled"`
	JSON    responseCollectionVrHfxGdBResultRuleSettingsPayloadLogJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsPayloadLogJSON contains the JSON
// metadata for the struct [ResponseCollectionVRHfxGdBResultRuleSettingsPayloadLog]
type responseCollectionVrHfxGdBResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid / an SSL error occurs.
type ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCert struct {
	// The action to perform upon seeing an untrusted certificate. Default action is
	// error with HTTP code 526.
	Action ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   responseCollectionVrHfxGdBResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// responseCollectionVrHfxGdBResultRuleSettingsUntrustedCertJSON contains the JSON
// metadata for the struct
// [ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCert]
type responseCollectionVrHfxGdBResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform upon seeing an untrusted certificate. Default action is
// error with HTTP code 526.
type ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertAction string

const (
	ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertActionPassThrough ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertAction = "pass_through"
	ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertActionBlock       ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertAction = "block"
	ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertActionError       ResponseCollectionVRHfxGdBResultRuleSettingsUntrustedCertAction = "error"
)

// Schedule for activating DNS policies. Does not apply to HTTP or network
// policies.
type ResponseCollectionVRHfxGdBResultSchedule struct {
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
	// parameter is omitted, then the time zone inferred from the user's source IP is
	// used to evaluate the rule. If Gateway cannot determine the time zone from the
	// IP, we will fall back to the time zone of the user's connected data center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string                                       `json:"wed"`
	JSON responseCollectionVrHfxGdBResultScheduleJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultScheduleJSON contains the JSON metadata for the
// struct [ResponseCollectionVRHfxGdBResultSchedule]
type responseCollectionVrHfxGdBResultScheduleJSON struct {
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

func (r *ResponseCollectionVRHfxGdBResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionVRHfxGdBResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionVrHfxGdBResultInfoJSON `json:"-"`
}

// responseCollectionVrHfxGdBResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionVRHfxGdBResultInfo]
type responseCollectionVrHfxGdBResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionVRHfxGdBResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionVRHfxGdBSuccess bool

const (
	ResponseCollectionVRHfxGdBSuccessTrue ResponseCollectionVRHfxGdBSuccess = true
)

type SingleResponseH8mJb2Ar struct {
	Errors   []SingleResponseH8mJb2ArError   `json:"errors"`
	Messages []SingleResponseH8mJb2ArMessage `json:"messages"`
	Result   SingleResponseH8mJb2ArResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseH8mJb2ArSuccess `json:"success"`
	JSON    singleResponseH8mJb2ArJSON    `json:"-"`
}

// singleResponseH8mJb2ArJSON contains the JSON metadata for the struct
// [SingleResponseH8mJb2Ar]
type singleResponseH8mJb2ArJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2Ar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseH8mJb2ArError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseH8mJb2ArErrorJSON `json:"-"`
}

// singleResponseH8mJb2ArErrorJSON contains the JSON metadata for the struct
// [SingleResponseH8mJb2ArError]
type singleResponseH8mJb2ArErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseH8mJb2ArMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseH8mJb2ArMessageJSON `json:"-"`
}

// singleResponseH8mJb2ArMessageJSON contains the JSON metadata for the struct
// [SingleResponseH8mJb2ArMessage]
type singleResponseH8mJb2ArMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseH8mJb2ArResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to 'true'.
	Action    SingleResponseH8mJb2ArResultAction `json:"action"`
	CreatedAt time.Time                          `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the Rule.
	Description string `json:"description"`
	// The wirefilter expression to be used for device posture check matching.
	DevicePosture string `json:"device_posture"`
	// Set if the rule is enabled.
	Enabled bool `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []SingleResponseH8mJb2ArResultFilter `json:"filters"`
	// The wirefilter expression to be used for identity matching.
	Identity string `json:"identity"`
	// The name of the Rule.
	Name string `json:"name"`
	// Precedence sets the ordering of the rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings SingleResponseH8mJb2ArResultRuleSettings `json:"rule_settings"`
	// Schedule for activating DNS policies. Does not apply to HTTP or network
	// policies.
	Schedule SingleResponseH8mJb2ArResultSchedule `json:"schedule"`
	// The wirefilter expression to be used for traffic matching.
	Traffic   string                           `json:"traffic"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      singleResponseH8mJb2ArResultJSON `json:"-"`
}

// singleResponseH8mJb2ArResultJSON contains the JSON metadata for the struct
// [SingleResponseH8mJb2ArResult]
type singleResponseH8mJb2ArResultJSON struct {
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

func (r *SingleResponseH8mJb2ArResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to 'true'.
type SingleResponseH8mJb2ArResultAction string

const (
	SingleResponseH8mJb2ArResultActionOn           SingleResponseH8mJb2ArResultAction = "on"
	SingleResponseH8mJb2ArResultActionOff          SingleResponseH8mJb2ArResultAction = "off"
	SingleResponseH8mJb2ArResultActionAllow        SingleResponseH8mJb2ArResultAction = "allow"
	SingleResponseH8mJb2ArResultActionBlock        SingleResponseH8mJb2ArResultAction = "block"
	SingleResponseH8mJb2ArResultActionScan         SingleResponseH8mJb2ArResultAction = "scan"
	SingleResponseH8mJb2ArResultActionNoscan       SingleResponseH8mJb2ArResultAction = "noscan"
	SingleResponseH8mJb2ArResultActionSafesearch   SingleResponseH8mJb2ArResultAction = "safesearch"
	SingleResponseH8mJb2ArResultActionYtrestricted SingleResponseH8mJb2ArResultAction = "ytrestricted"
	SingleResponseH8mJb2ArResultActionIsolate      SingleResponseH8mJb2ArResultAction = "isolate"
	SingleResponseH8mJb2ArResultActionNoisolate    SingleResponseH8mJb2ArResultAction = "noisolate"
	SingleResponseH8mJb2ArResultActionOverride     SingleResponseH8mJb2ArResultAction = "override"
	SingleResponseH8mJb2ArResultActionL4Override   SingleResponseH8mJb2ArResultAction = "l4_override"
	SingleResponseH8mJb2ArResultActionEgress       SingleResponseH8mJb2ArResultAction = "egress"
	SingleResponseH8mJb2ArResultActionAuditSSH     SingleResponseH8mJb2ArResultAction = "audit_ssh"
)

// The protocol or layer to use.
type SingleResponseH8mJb2ArResultFilter string

const (
	SingleResponseH8mJb2ArResultFilterHTTP   SingleResponseH8mJb2ArResultFilter = "http"
	SingleResponseH8mJb2ArResultFilterDNS    SingleResponseH8mJb2ArResultFilter = "dns"
	SingleResponseH8mJb2ArResultFilterL4     SingleResponseH8mJb2ArResultFilter = "l4"
	SingleResponseH8mJb2ArResultFilterEgress SingleResponseH8mJb2ArResultFilter = "egress"
)

// Additional settings that modify the rule's action.
type SingleResponseH8mJb2ArResultRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Audit ssh action settings
	AuditSSH SingleResponseH8mJb2ArResultRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls SingleResponseH8mJb2ArResultRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred that will be displayed on the custom
	// block page (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession SingleResponseH8mJb2ArResultRuleSettingsCheckSession `json:"check_session"`
	// Configure how Proxy traffic egresses. Can be set for rules with Egress action
	// and Egress filter. Can be omitted to indicate local egress via Warp IPs
	Egress SingleResponseH8mJb2ArResultRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Include IPs in DNS resolver category blocks. By default categories only block on
	// domain names.
	IPCategories bool `json:"ip_categories"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override SingleResponseH8mJb2ArResultRuleSettingsL4override `json:"l4override"`
	// Override matching DNS queries with this.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with this.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog SingleResponseH8mJb2ArResultRuleSettingsPayloadLog `json:"payload_log"`
	// Configure behavior when an upstream cert is invalid / an SSL error occurs.
	UntrustedCert SingleResponseH8mJb2ArResultRuleSettingsUntrustedCert `json:"untrusted_cert"`
	JSON          singleResponseH8mJb2ArResultRuleSettingsJSON          `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsJSON contains the JSON metadata for the
// struct [SingleResponseH8mJb2ArResultRuleSettings]
type singleResponseH8mJb2ArResultRuleSettingsJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BisoAdminControls               apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	Egress                          apijson.Field
	InsecureDisableDnssecValidation apijson.Field
	IPCategories                    apijson.Field
	L4override                      apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	UntrustedCert                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Audit ssh action settings
type SingleResponseH8mJb2ArResultRuleSettingsAuditSSH struct {
	// Turn on SSH command logging.
	CommandLogging bool                                                 `json:"command_logging"`
	JSON           singleResponseH8mJb2ArResultRuleSettingsAuditSSHJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsAuditSSHJSON contains the JSON metadata
// for the struct [SingleResponseH8mJb2ArResultRuleSettingsAuditSSH]
type singleResponseH8mJb2ArResultRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how browser isolation behaves.
type SingleResponseH8mJb2ArResultRuleSettingsBisoAdminControls struct {
	// Disable copy-paste.
	Dcp bool `json:"dcp"`
	// Disable download.
	Dd bool `json:"dd"`
	// Disable keyboard usage.
	Dk bool `json:"dk"`
	// Disable printing.
	Dp bool `json:"dp"`
	// Disable upload.
	Du   bool                                                          `json:"du"`
	JSON singleResponseH8mJb2ArResultRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct
// [SingleResponseH8mJb2ArResultRuleSettingsBisoAdminControls]
type singleResponseH8mJb2ArResultRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how session check behaves.
type SingleResponseH8mJb2ArResultRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Enable session enforcement for this fule.
	Enforce bool                                                     `json:"enforce"`
	JSON    singleResponseH8mJb2ArResultRuleSettingsCheckSessionJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsCheckSessionJSON contains the JSON
// metadata for the struct [SingleResponseH8mJb2ArResultRuleSettingsCheckSession]
type singleResponseH8mJb2ArResultRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure how Proxy traffic egresses. Can be set for rules with Egress action
// and Egress filter. Can be omitted to indicate local egress via Warp IPs
type SingleResponseH8mJb2ArResultRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The IPv4 address to be used for egress in the event of an error egressing with
	// the primary IPv4. Can be '0.0.0.0' to indicate local egreass via Warp IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                                             `json:"ipv6"`
	JSON singleResponseH8mJb2ArResultRuleSettingsEgressJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsEgressJSON contains the JSON metadata
// for the struct [SingleResponseH8mJb2ArResultRuleSettingsEgress]
type singleResponseH8mJb2ArResultRuleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Send matching traffic to the supplied destination IP address and port.
type SingleResponseH8mJb2ArResultRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                                  `json:"port"`
	JSON singleResponseH8mJb2ArResultRuleSettingsL4overrideJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsL4overrideJSON contains the JSON
// metadata for the struct [SingleResponseH8mJb2ArResultRuleSettingsL4override]
type singleResponseH8mJb2ArResultRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure DLP payload logging.
type SingleResponseH8mJb2ArResultRuleSettingsPayloadLog struct {
	// Enable DLP payload logging for this rule.
	Enabled bool                                                   `json:"enabled"`
	JSON    singleResponseH8mJb2ArResultRuleSettingsPayloadLogJSON `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsPayloadLogJSON contains the JSON
// metadata for the struct [SingleResponseH8mJb2ArResultRuleSettingsPayloadLog]
type singleResponseH8mJb2ArResultRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure behavior when an upstream cert is invalid / an SSL error occurs.
type SingleResponseH8mJb2ArResultRuleSettingsUntrustedCert struct {
	// The action to perform upon seeing an untrusted certificate. Default action is
	// error with HTTP code 526.
	Action SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertAction `json:"action"`
	JSON   singleResponseH8mJb2ArResultRuleSettingsUntrustedCertJSON   `json:"-"`
}

// singleResponseH8mJb2ArResultRuleSettingsUntrustedCertJSON contains the JSON
// metadata for the struct [SingleResponseH8mJb2ArResultRuleSettingsUntrustedCert]
type singleResponseH8mJb2ArResultRuleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseH8mJb2ArResultRuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform upon seeing an untrusted certificate. Default action is
// error with HTTP code 526.
type SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertAction string

const (
	SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertActionPassThrough SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertAction = "pass_through"
	SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertActionBlock       SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertAction = "block"
	SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertActionError       SingleResponseH8mJb2ArResultRuleSettingsUntrustedCertAction = "error"
)

// Schedule for activating DNS policies. Does not apply to HTTP or network
// policies.
type SingleResponseH8mJb2ArResultSchedule struct {
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
	// parameter is omitted, then the time zone inferred from the user's source IP is
	// used to evaluate the rule. If Gateway cannot determine the time zone from the
	// IP, we will fall back to the time zone of the user's connected data center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string                                   `json:"wed"`
	JSON singleResponseH8mJb2ArResultScheduleJSON `json:"-"`
}

// singleResponseH8mJb2ArResultScheduleJSON contains the JSON metadata for the
// struct [SingleResponseH8mJb2ArResultSchedule]
type singleResponseH8mJb2ArResultScheduleJSON struct {
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

func (r *SingleResponseH8mJb2ArResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseH8mJb2ArSuccess bool

const (
	SingleResponseH8mJb2ArSuccessTrue SingleResponseH8mJb2ArSuccess = true
)

type AccountGatewayRuleUpdateParams struct {
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to 'true'.
	Action param.Field[AccountGatewayRuleUpdateParamsAction] `json:"action,required"`
	// The name of the Rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the Rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression to be used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// Set if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]AccountGatewayRuleUpdateParamsFilter] `json:"filters"`
	// The wirefilter expression to be used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the ordering of the rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[AccountGatewayRuleUpdateParamsRuleSettings] `json:"rule_settings"`
	// Schedule for activating DNS policies. Does not apply to HTTP or network
	// policies.
	Schedule param.Field[AccountGatewayRuleUpdateParamsSchedule] `json:"schedule"`
	// The wirefilter expression to be used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to 'true'.
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
	// Audit ssh action settings
	AuditSSH param.Field[AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred that will be displayed on the custom
	// block page (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[AccountGatewayRuleUpdateParamsRuleSettingsCheckSession] `json:"check_session"`
	// Configure how Proxy traffic egresses. Can be set for rules with Egress action
	// and Egress filter. Can be omitted to indicate local egress via Warp IPs
	Egress param.Field[AccountGatewayRuleUpdateParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for allow actions).
	InsecureDisableDnssecValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Include IPs in DNS resolver category blocks. By default categories only block on
	// domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[AccountGatewayRuleUpdateParamsRuleSettingsL4override] `json:"l4override"`
	// Override matching DNS queries with this.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with this.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Configure behavior when an upstream cert is invalid / an SSL error occurs.
	UntrustedCert param.Field[AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audit ssh action settings
type AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH struct {
	// Turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls struct {
	// Disable copy-paste.
	Dcp param.Field[bool] `json:"dcp"`
	// Disable download.
	Dd param.Field[bool] `json:"dd"`
	// Disable keyboard usage.
	Dk param.Field[bool] `json:"dk"`
	// Disable printing.
	Dp param.Field[bool] `json:"dp"`
	// Disable upload.
	Du param.Field[bool] `json:"du"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type AccountGatewayRuleUpdateParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Enable session enforcement for this fule.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Proxy traffic egresses. Can be set for rules with Egress action
// and Egress filter. Can be omitted to indicate local egress via Warp IPs
type AccountGatewayRuleUpdateParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 param.Field[string] `json:"ipv4"`
	// The IPv4 address to be used for egress in the event of an error egressing with
	// the primary IPv4. Can be '0.0.0.0' to indicate local egreass via Warp IPs.
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

// Configure DLP payload logging.
type AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog struct {
	// Enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid / an SSL error occurs.
type AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert struct {
	// The action to perform upon seeing an untrusted certificate. Default action is
	// error with HTTP code 526.
	Action param.Field[AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform upon seeing an untrusted certificate. Default action is
// error with HTTP code 526.
type AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertActionError       AccountGatewayRuleUpdateParamsRuleSettingsUntrustedCertAction = "error"
)

// Schedule for activating DNS policies. Does not apply to HTTP or network
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
	// parameter is omitted, then the time zone inferred from the user's source IP is
	// used to evaluate the rule. If Gateway cannot determine the time zone from the
	// IP, we will fall back to the time zone of the user's connected data center.
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
	// expressions are either absent or evaluate to 'true'.
	Action param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsAction] `json:"action,required"`
	// The name of the Rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the Rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression to be used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// Set if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsFilter] `json:"filters"`
	// The wirefilter expression to be used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the ordering of the rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings] `json:"rule_settings"`
	// Schedule for activating DNS policies. Does not apply to HTTP or network
	// policies.
	Schedule param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsSchedule] `json:"schedule"`
	// The wirefilter expression to be used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to 'true'.
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
	// Audit ssh action settings
	AuditSSH param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred that will be displayed on the custom
	// block page (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession] `json:"check_session"`
	// Configure how Proxy traffic egresses. Can be set for rules with Egress action
	// and Egress filter. Can be omitted to indicate local egress via Warp IPs
	Egress param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for allow actions).
	InsecureDisableDnssecValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Include IPs in DNS resolver category blocks. By default categories only block on
	// domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsL4override] `json:"l4override"`
	// Override matching DNS queries with this.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with this.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Configure behavior when an upstream cert is invalid / an SSL error occurs.
	UntrustedCert param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert] `json:"untrusted_cert"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audit ssh action settings
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH struct {
	// Turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls struct {
	// Disable copy-paste.
	Dcp param.Field[bool] `json:"dcp"`
	// Disable download.
	Dd param.Field[bool] `json:"dd"`
	// Disable keyboard usage.
	Dk param.Field[bool] `json:"dk"`
	// Disable printing.
	Dp param.Field[bool] `json:"dp"`
	// Disable upload.
	Du param.Field[bool] `json:"du"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Enable session enforcement for this fule.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Proxy traffic egresses. Can be set for rules with Egress action
// and Egress filter. Can be omitted to indicate local egress via Warp IPs
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 param.Field[string] `json:"ipv4"`
	// The IPv4 address to be used for egress in the event of an error egressing with
	// the primary IPv4. Can be '0.0.0.0' to indicate local egreass via Warp IPs.
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

// Configure DLP payload logging.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog struct {
	// Enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid / an SSL error occurs.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert struct {
	// The action to perform upon seeing an untrusted certificate. Default action is
	// error with HTTP code 526.
	Action param.Field[AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction] `json:"action"`
}

func (r AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCert) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform upon seeing an untrusted certificate. Default action is
// error with HTTP code 526.
type AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction string

const (
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionPassThrough AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "pass_through"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionBlock       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "block"
	AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertActionError       AccountGatewayRuleZeroTrustGatewayRulesNewZeroTrustGatewayRuleParamsRuleSettingsUntrustedCertAction = "error"
)

// Schedule for activating DNS policies. Does not apply to HTTP or network
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
	// parameter is omitted, then the time zone inferred from the user's source IP is
	// used to evaluate the rule. If Gateway cannot determine the time zone from the
	// IP, we will fall back to the time zone of the user's connected data center.
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

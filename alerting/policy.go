// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

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

// PolicyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r *PolicyService) {
	r = &PolicyService{}
	r.Options = opts
	return
}

// Creates a new Notification policy.
func (r *PolicyService) New(ctx context.Context, params PolicyNewParams, opts ...option.RequestOption) (res *PolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Notification policy.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *PolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PolicyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all Notification policies.
func (r *PolicyService) List(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AlertingPolicies], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", query.AccountID)
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

// Get a list of all Notification policies.
func (r *PolicyService) ListAutoPaging(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AlertingPolicies] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Notification policy.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (res *PolicyDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single policy.
func (r *PolicyService) Get(ctx context.Context, policyID string, query PolicyGetParams, opts ...option.RequestOption) (res *AlertingPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env PolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingPolicies struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertingPoliciesAlertType `json:"alert_type"`
	Created   time.Time                 `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertingPoliciesFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]AlertingPoliciesMechanisms `json:"mechanisms"`
	Modified   time.Time                               `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string               `json:"name"`
	JSON alertingPoliciesJSON `json:"-"`
}

// alertingPoliciesJSON contains the JSON metadata for the struct
// [AlertingPolicies]
type alertingPoliciesJSON struct {
	ID          apijson.Field
	AlertType   apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Filters     apijson.Field
	Mechanisms  apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertingPoliciesJSON) RawJSON() string {
	return r.raw
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingPoliciesAlertType string

const (
	AlertingPoliciesAlertTypeAccessCustomCertificateExpirationType         AlertingPoliciesAlertType = "access_custom_certificate_expiration_type"
	AlertingPoliciesAlertTypeAdvancedDDoSAttackL4Alert                     AlertingPoliciesAlertType = "advanced_ddos_attack_l4_alert"
	AlertingPoliciesAlertTypeAdvancedDDoSAttackL7Alert                     AlertingPoliciesAlertType = "advanced_ddos_attack_l7_alert"
	AlertingPoliciesAlertTypeAdvancedHTTPAlertError                        AlertingPoliciesAlertType = "advanced_http_alert_error"
	AlertingPoliciesAlertTypeBGPHijackNotification                         AlertingPoliciesAlertType = "bgp_hijack_notification"
	AlertingPoliciesAlertTypeBillingUsageAlert                             AlertingPoliciesAlertType = "billing_usage_alert"
	AlertingPoliciesAlertTypeBlockNotificationBlockRemoved                 AlertingPoliciesAlertType = "block_notification_block_removed"
	AlertingPoliciesAlertTypeBlockNotificationNewBlock                     AlertingPoliciesAlertType = "block_notification_new_block"
	AlertingPoliciesAlertTypeBlockNotificationReviewRejected               AlertingPoliciesAlertType = "block_notification_review_rejected"
	AlertingPoliciesAlertTypeBrandProtectionAlert                          AlertingPoliciesAlertType = "brand_protection_alert"
	AlertingPoliciesAlertTypeBrandProtectionDigest                         AlertingPoliciesAlertType = "brand_protection_digest"
	AlertingPoliciesAlertTypeClickhouseAlertFwAnomaly                      AlertingPoliciesAlertType = "clickhouse_alert_fw_anomaly"
	AlertingPoliciesAlertTypeClickhouseAlertFwEntAnomaly                   AlertingPoliciesAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingPoliciesAlertTypeCustomSSLCertificateEventType                 AlertingPoliciesAlertType = "custom_ssl_certificate_event_type"
	AlertingPoliciesAlertTypeDedicatedSSLCertificateEventType              AlertingPoliciesAlertType = "dedicated_ssl_certificate_event_type"
	AlertingPoliciesAlertTypeDosAttackL4                                   AlertingPoliciesAlertType = "dos_attack_l4"
	AlertingPoliciesAlertTypeDosAttackL7                                   AlertingPoliciesAlertType = "dos_attack_l7"
	AlertingPoliciesAlertTypeExpiringServiceTokenAlert                     AlertingPoliciesAlertType = "expiring_service_token_alert"
	AlertingPoliciesAlertTypeFailingLogpushJobDisabledAlert                AlertingPoliciesAlertType = "failing_logpush_job_disabled_alert"
	AlertingPoliciesAlertTypeFbmAutoAdvertisement                          AlertingPoliciesAlertType = "fbm_auto_advertisement"
	AlertingPoliciesAlertTypeFbmDosdAttack                                 AlertingPoliciesAlertType = "fbm_dosd_attack"
	AlertingPoliciesAlertTypeFbmVolumetricAttack                           AlertingPoliciesAlertType = "fbm_volumetric_attack"
	AlertingPoliciesAlertTypeHealthCheckStatusNotification                 AlertingPoliciesAlertType = "health_check_status_notification"
	AlertingPoliciesAlertTypeHostnameAopCustomCertificateExpirationType    AlertingPoliciesAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingPoliciesAlertTypeHTTPAlertEdgeError                            AlertingPoliciesAlertType = "http_alert_edge_error"
	AlertingPoliciesAlertTypeHTTPAlertOriginError                          AlertingPoliciesAlertType = "http_alert_origin_error"
	AlertingPoliciesAlertTypeIncidentAlert                                 AlertingPoliciesAlertType = "incident_alert"
	AlertingPoliciesAlertTypeLoadBalancingHealthAlert                      AlertingPoliciesAlertType = "load_balancing_health_alert"
	AlertingPoliciesAlertTypeLoadBalancingPoolEnablementAlert              AlertingPoliciesAlertType = "load_balancing_pool_enablement_alert"
	AlertingPoliciesAlertTypeLogoMatchAlert                                AlertingPoliciesAlertType = "logo_match_alert"
	AlertingPoliciesAlertTypeMagicTunnelHealthCheckEvent                   AlertingPoliciesAlertType = "magic_tunnel_health_check_event"
	AlertingPoliciesAlertTypeMaintenanceEventNotification                  AlertingPoliciesAlertType = "maintenance_event_notification"
	AlertingPoliciesAlertTypeMTLSCertificateStoreCertificateExpirationType AlertingPoliciesAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingPoliciesAlertTypePagesEventAlert                               AlertingPoliciesAlertType = "pages_event_alert"
	AlertingPoliciesAlertTypeRadarNotification                             AlertingPoliciesAlertType = "radar_notification"
	AlertingPoliciesAlertTypeRealOriginMonitoring                          AlertingPoliciesAlertType = "real_origin_monitoring"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingPoliciesAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewHosts                    AlertingPoliciesAlertType = "scriptmonitor_alert_new_hosts"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingPoliciesAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingPoliciesAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingPoliciesAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingPoliciesAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingPoliciesAlertTypeScriptmonitorAlertNewResources                AlertingPoliciesAlertType = "scriptmonitor_alert_new_resources"
	AlertingPoliciesAlertTypeSecondaryDNSAllPrimariesFailing               AlertingPoliciesAlertType = "secondary_dns_all_primaries_failing"
	AlertingPoliciesAlertTypeSecondaryDNSPrimariesFailing                  AlertingPoliciesAlertType = "secondary_dns_primaries_failing"
	AlertingPoliciesAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingPoliciesAlertType = "secondary_dns_zone_successfully_updated"
	AlertingPoliciesAlertTypeSecondaryDNSZoneValidationWarning             AlertingPoliciesAlertType = "secondary_dns_zone_validation_warning"
	AlertingPoliciesAlertTypeSentinelAlert                                 AlertingPoliciesAlertType = "sentinel_alert"
	AlertingPoliciesAlertTypeStreamLiveNotifications                       AlertingPoliciesAlertType = "stream_live_notifications"
	AlertingPoliciesAlertTypeTrafficAnomaliesAlert                         AlertingPoliciesAlertType = "traffic_anomalies_alert"
	AlertingPoliciesAlertTypeTunnelHealthEvent                             AlertingPoliciesAlertType = "tunnel_health_event"
	AlertingPoliciesAlertTypeTunnelUpdateEvent                             AlertingPoliciesAlertType = "tunnel_update_event"
	AlertingPoliciesAlertTypeUniversalSSLEventType                         AlertingPoliciesAlertType = "universal_ssl_event_type"
	AlertingPoliciesAlertTypeWebAnalyticsMetricsUpdate                     AlertingPoliciesAlertType = "web_analytics_metrics_update"
	AlertingPoliciesAlertTypeZoneAopCustomCertificateExpirationType        AlertingPoliciesAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r AlertingPoliciesAlertType) IsKnown() bool {
	switch r {
	case AlertingPoliciesAlertTypeAccessCustomCertificateExpirationType, AlertingPoliciesAlertTypeAdvancedDDoSAttackL4Alert, AlertingPoliciesAlertTypeAdvancedDDoSAttackL7Alert, AlertingPoliciesAlertTypeAdvancedHTTPAlertError, AlertingPoliciesAlertTypeBGPHijackNotification, AlertingPoliciesAlertTypeBillingUsageAlert, AlertingPoliciesAlertTypeBlockNotificationBlockRemoved, AlertingPoliciesAlertTypeBlockNotificationNewBlock, AlertingPoliciesAlertTypeBlockNotificationReviewRejected, AlertingPoliciesAlertTypeBrandProtectionAlert, AlertingPoliciesAlertTypeBrandProtectionDigest, AlertingPoliciesAlertTypeClickhouseAlertFwAnomaly, AlertingPoliciesAlertTypeClickhouseAlertFwEntAnomaly, AlertingPoliciesAlertTypeCustomSSLCertificateEventType, AlertingPoliciesAlertTypeDedicatedSSLCertificateEventType, AlertingPoliciesAlertTypeDosAttackL4, AlertingPoliciesAlertTypeDosAttackL7, AlertingPoliciesAlertTypeExpiringServiceTokenAlert, AlertingPoliciesAlertTypeFailingLogpushJobDisabledAlert, AlertingPoliciesAlertTypeFbmAutoAdvertisement, AlertingPoliciesAlertTypeFbmDosdAttack, AlertingPoliciesAlertTypeFbmVolumetricAttack, AlertingPoliciesAlertTypeHealthCheckStatusNotification, AlertingPoliciesAlertTypeHostnameAopCustomCertificateExpirationType, AlertingPoliciesAlertTypeHTTPAlertEdgeError, AlertingPoliciesAlertTypeHTTPAlertOriginError, AlertingPoliciesAlertTypeIncidentAlert, AlertingPoliciesAlertTypeLoadBalancingHealthAlert, AlertingPoliciesAlertTypeLoadBalancingPoolEnablementAlert, AlertingPoliciesAlertTypeLogoMatchAlert, AlertingPoliciesAlertTypeMagicTunnelHealthCheckEvent, AlertingPoliciesAlertTypeMaintenanceEventNotification, AlertingPoliciesAlertTypeMTLSCertificateStoreCertificateExpirationType, AlertingPoliciesAlertTypePagesEventAlert, AlertingPoliciesAlertTypeRadarNotification, AlertingPoliciesAlertTypeRealOriginMonitoring, AlertingPoliciesAlertTypeScriptmonitorAlertNewCodeChangeDetections, AlertingPoliciesAlertTypeScriptmonitorAlertNewHosts, AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousHosts, AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousScripts, AlertingPoliciesAlertTypeScriptmonitorAlertNewMaliciousURL, AlertingPoliciesAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, AlertingPoliciesAlertTypeScriptmonitorAlertNewResources, AlertingPoliciesAlertTypeSecondaryDNSAllPrimariesFailing, AlertingPoliciesAlertTypeSecondaryDNSPrimariesFailing, AlertingPoliciesAlertTypeSecondaryDNSZoneSuccessfullyUpdated, AlertingPoliciesAlertTypeSecondaryDNSZoneValidationWarning, AlertingPoliciesAlertTypeSentinelAlert, AlertingPoliciesAlertTypeStreamLiveNotifications, AlertingPoliciesAlertTypeTrafficAnomaliesAlert, AlertingPoliciesAlertTypeTunnelHealthEvent, AlertingPoliciesAlertTypeTunnelUpdateEvent, AlertingPoliciesAlertTypeUniversalSSLEventType, AlertingPoliciesAlertTypeWebAnalyticsMetricsUpdate, AlertingPoliciesAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingPoliciesFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs []string `json:"affected_asns"`
	// Used for configuring incident_alert. A list of identifiers for each component to
	// monitor.
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue []AlertingPoliciesFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled []string `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment []string `json:"environment"`
	// Used for configuring pages_event_alert
	Event []string `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource []string `json:"event_source"`
	// Usage depends on specific alert type
	EventType []string `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy []string `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID []string `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact []AlertingPoliciesFiltersIncidentImpact `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID []string `json:"input_id"`
	// Used for configuring billing_usage_alert
	Limit []string `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag []string `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond []string `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth []string `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus []string `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond []string `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID []string `json:"pool_id"`
	// Used for configuring billing_usage_alert
	Product []string `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID []string `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol []string `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag []string `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond []string `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors []string `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services []string `json:"services"`
	// Usage depends on specific alert type
	Slo []string `json:"slo"`
	// Used for configuring health_check_status_notification
	Status []string `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname []string `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP []string `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName []string `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions []AlertingPoliciesFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                    `json:"zones"`
	JSON  alertingPoliciesFiltersJSON `json:"-"`
}

// alertingPoliciesFiltersJSON contains the JSON metadata for the struct
// [AlertingPoliciesFilters]
type alertingPoliciesFiltersJSON struct {
	Actions                      apijson.Field
	AffectedASNs                 apijson.Field
	AffectedComponents           apijson.Field
	AffectedLocations            apijson.Field
	AirportCode                  apijson.Field
	AlertTriggerPreferences      apijson.Field
	AlertTriggerPreferencesValue apijson.Field
	Enabled                      apijson.Field
	Environment                  apijson.Field
	Event                        apijson.Field
	EventSource                  apijson.Field
	EventType                    apijson.Field
	GroupBy                      apijson.Field
	HealthCheckID                apijson.Field
	IncidentImpact               apijson.Field
	InputID                      apijson.Field
	Limit                        apijson.Field
	LogoTag                      apijson.Field
	MegabitsPerSecond            apijson.Field
	NewHealth                    apijson.Field
	NewStatus                    apijson.Field
	PacketsPerSecond             apijson.Field
	PoolID                       apijson.Field
	Product                      apijson.Field
	ProjectID                    apijson.Field
	Protocol                     apijson.Field
	QueryTag                     apijson.Field
	RequestsPerSecond            apijson.Field
	Selectors                    apijson.Field
	Services                     apijson.Field
	Slo                          apijson.Field
	Status                       apijson.Field
	TargetHostname               apijson.Field
	TargetIP                     apijson.Field
	TargetZoneName               apijson.Field
	TrafficExclusions            apijson.Field
	TunnelID                     apijson.Field
	TunnelName                   apijson.Field
	Where                        apijson.Field
	Zones                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AlertingPoliciesFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertingPoliciesFiltersJSON) RawJSON() string {
	return r.raw
}

type AlertingPoliciesFiltersAlertTriggerPreferencesValue string

const (
	AlertingPoliciesFiltersAlertTriggerPreferencesValue99_0 AlertingPoliciesFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingPoliciesFiltersAlertTriggerPreferencesValue98_0 AlertingPoliciesFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingPoliciesFiltersAlertTriggerPreferencesValue97_0 AlertingPoliciesFiltersAlertTriggerPreferencesValue = "97.0"
)

func (r AlertingPoliciesFiltersAlertTriggerPreferencesValue) IsKnown() bool {
	switch r {
	case AlertingPoliciesFiltersAlertTriggerPreferencesValue99_0, AlertingPoliciesFiltersAlertTriggerPreferencesValue98_0, AlertingPoliciesFiltersAlertTriggerPreferencesValue97_0:
		return true
	}
	return false
}

type AlertingPoliciesFiltersIncidentImpact string

const (
	AlertingPoliciesFiltersIncidentImpactIncidentImpactNone     AlertingPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingPoliciesFiltersIncidentImpactIncidentImpactMinor    AlertingPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingPoliciesFiltersIncidentImpactIncidentImpactMajor    AlertingPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingPoliciesFiltersIncidentImpactIncidentImpactCritical AlertingPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

func (r AlertingPoliciesFiltersIncidentImpact) IsKnown() bool {
	switch r {
	case AlertingPoliciesFiltersIncidentImpactIncidentImpactNone, AlertingPoliciesFiltersIncidentImpactIncidentImpactMinor, AlertingPoliciesFiltersIncidentImpactIncidentImpactMajor, AlertingPoliciesFiltersIncidentImpactIncidentImpactCritical:
		return true
	}
	return false
}

type AlertingPoliciesFiltersTrafficExclusion string

const (
	AlertingPoliciesFiltersTrafficExclusionSecurityEvents AlertingPoliciesFiltersTrafficExclusion = "security_events"
)

func (r AlertingPoliciesFiltersTrafficExclusion) IsKnown() bool {
	switch r {
	case AlertingPoliciesFiltersTrafficExclusionSecurityEvents:
		return true
	}
	return false
}

type AlertingPoliciesMechanisms struct {
	// UUID
	ID   AlertingPoliciesMechanismsIDUnion `json:"id"`
	JSON alertingPoliciesMechanismsJSON    `json:"-"`
}

// alertingPoliciesMechanismsJSON contains the JSON metadata for the struct
// [AlertingPoliciesMechanisms]
type alertingPoliciesMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingPoliciesMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertingPoliciesMechanismsJSON) RawJSON() string {
	return r.raw
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AlertingPoliciesMechanismsIDUnion interface {
	ImplementsAlertingAlertingPoliciesMechanismsIDUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingPoliciesMechanismsIDUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PolicyNewResponse struct {
	// UUID
	ID   string                `json:"id"`
	JSON policyNewResponseJSON `json:"-"`
}

// policyNewResponseJSON contains the JSON metadata for the struct
// [PolicyNewResponse]
type policyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseJSON) RawJSON() string {
	return r.raw
}

type PolicyUpdateResponse struct {
	// UUID
	ID   string                   `json:"id"`
	JSON policyUpdateResponseJSON `json:"-"`
}

// policyUpdateResponseJSON contains the JSON metadata for the struct
// [PolicyUpdateResponse]
type policyUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [alerting.PolicyDeleteResponseUnknown],
// [alerting.PolicyDeleteResponseArray] or [shared.UnionString].
type PolicyDeleteResponseUnion interface {
	ImplementsAlertingPolicyDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PolicyDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PolicyDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PolicyDeleteResponseArray []interface{}

func (r PolicyDeleteResponseArray) ImplementsAlertingPolicyDeleteResponseUnion() {}

type PolicyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[PolicyNewParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]PolicyNewParamsMechanisms] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[PolicyNewParamsFilters] `json:"filters"`
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type PolicyNewParamsAlertType string

const (
	PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType         PolicyNewParamsAlertType = "access_custom_certificate_expiration_type"
	PolicyNewParamsAlertTypeAdvancedDDoSAttackL4Alert                     PolicyNewParamsAlertType = "advanced_ddos_attack_l4_alert"
	PolicyNewParamsAlertTypeAdvancedDDoSAttackL7Alert                     PolicyNewParamsAlertType = "advanced_ddos_attack_l7_alert"
	PolicyNewParamsAlertTypeAdvancedHTTPAlertError                        PolicyNewParamsAlertType = "advanced_http_alert_error"
	PolicyNewParamsAlertTypeBGPHijackNotification                         PolicyNewParamsAlertType = "bgp_hijack_notification"
	PolicyNewParamsAlertTypeBillingUsageAlert                             PolicyNewParamsAlertType = "billing_usage_alert"
	PolicyNewParamsAlertTypeBlockNotificationBlockRemoved                 PolicyNewParamsAlertType = "block_notification_block_removed"
	PolicyNewParamsAlertTypeBlockNotificationNewBlock                     PolicyNewParamsAlertType = "block_notification_new_block"
	PolicyNewParamsAlertTypeBlockNotificationReviewRejected               PolicyNewParamsAlertType = "block_notification_review_rejected"
	PolicyNewParamsAlertTypeBrandProtectionAlert                          PolicyNewParamsAlertType = "brand_protection_alert"
	PolicyNewParamsAlertTypeBrandProtectionDigest                         PolicyNewParamsAlertType = "brand_protection_digest"
	PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly                      PolicyNewParamsAlertType = "clickhouse_alert_fw_anomaly"
	PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly                   PolicyNewParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	PolicyNewParamsAlertTypeCustomSSLCertificateEventType                 PolicyNewParamsAlertType = "custom_ssl_certificate_event_type"
	PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType              PolicyNewParamsAlertType = "dedicated_ssl_certificate_event_type"
	PolicyNewParamsAlertTypeDosAttackL4                                   PolicyNewParamsAlertType = "dos_attack_l4"
	PolicyNewParamsAlertTypeDosAttackL7                                   PolicyNewParamsAlertType = "dos_attack_l7"
	PolicyNewParamsAlertTypeExpiringServiceTokenAlert                     PolicyNewParamsAlertType = "expiring_service_token_alert"
	PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert                PolicyNewParamsAlertType = "failing_logpush_job_disabled_alert"
	PolicyNewParamsAlertTypeFbmAutoAdvertisement                          PolicyNewParamsAlertType = "fbm_auto_advertisement"
	PolicyNewParamsAlertTypeFbmDosdAttack                                 PolicyNewParamsAlertType = "fbm_dosd_attack"
	PolicyNewParamsAlertTypeFbmVolumetricAttack                           PolicyNewParamsAlertType = "fbm_volumetric_attack"
	PolicyNewParamsAlertTypeHealthCheckStatusNotification                 PolicyNewParamsAlertType = "health_check_status_notification"
	PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType    PolicyNewParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	PolicyNewParamsAlertTypeHTTPAlertEdgeError                            PolicyNewParamsAlertType = "http_alert_edge_error"
	PolicyNewParamsAlertTypeHTTPAlertOriginError                          PolicyNewParamsAlertType = "http_alert_origin_error"
	PolicyNewParamsAlertTypeIncidentAlert                                 PolicyNewParamsAlertType = "incident_alert"
	PolicyNewParamsAlertTypeLoadBalancingHealthAlert                      PolicyNewParamsAlertType = "load_balancing_health_alert"
	PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert              PolicyNewParamsAlertType = "load_balancing_pool_enablement_alert"
	PolicyNewParamsAlertTypeLogoMatchAlert                                PolicyNewParamsAlertType = "logo_match_alert"
	PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent                   PolicyNewParamsAlertType = "magic_tunnel_health_check_event"
	PolicyNewParamsAlertTypeMaintenanceEventNotification                  PolicyNewParamsAlertType = "maintenance_event_notification"
	PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType PolicyNewParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	PolicyNewParamsAlertTypePagesEventAlert                               PolicyNewParamsAlertType = "pages_event_alert"
	PolicyNewParamsAlertTypeRadarNotification                             PolicyNewParamsAlertType = "radar_notification"
	PolicyNewParamsAlertTypeRealOriginMonitoring                          PolicyNewParamsAlertType = "real_origin_monitoring"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     PolicyNewParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts                    PolicyNewParamsAlertType = "scriptmonitor_alert_new_hosts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL             PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     PolicyNewParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewResources                PolicyNewParamsAlertType = "scriptmonitor_alert_new_resources"
	PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing               PolicyNewParamsAlertType = "secondary_dns_all_primaries_failing"
	PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing                  PolicyNewParamsAlertType = "secondary_dns_primaries_failing"
	PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           PolicyNewParamsAlertType = "secondary_dns_zone_successfully_updated"
	PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning             PolicyNewParamsAlertType = "secondary_dns_zone_validation_warning"
	PolicyNewParamsAlertTypeSentinelAlert                                 PolicyNewParamsAlertType = "sentinel_alert"
	PolicyNewParamsAlertTypeStreamLiveNotifications                       PolicyNewParamsAlertType = "stream_live_notifications"
	PolicyNewParamsAlertTypeTrafficAnomaliesAlert                         PolicyNewParamsAlertType = "traffic_anomalies_alert"
	PolicyNewParamsAlertTypeTunnelHealthEvent                             PolicyNewParamsAlertType = "tunnel_health_event"
	PolicyNewParamsAlertTypeTunnelUpdateEvent                             PolicyNewParamsAlertType = "tunnel_update_event"
	PolicyNewParamsAlertTypeUniversalSSLEventType                         PolicyNewParamsAlertType = "universal_ssl_event_type"
	PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate                     PolicyNewParamsAlertType = "web_analytics_metrics_update"
	PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType        PolicyNewParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r PolicyNewParamsAlertType) IsKnown() bool {
	switch r {
	case PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType, PolicyNewParamsAlertTypeAdvancedDDoSAttackL4Alert, PolicyNewParamsAlertTypeAdvancedDDoSAttackL7Alert, PolicyNewParamsAlertTypeAdvancedHTTPAlertError, PolicyNewParamsAlertTypeBGPHijackNotification, PolicyNewParamsAlertTypeBillingUsageAlert, PolicyNewParamsAlertTypeBlockNotificationBlockRemoved, PolicyNewParamsAlertTypeBlockNotificationNewBlock, PolicyNewParamsAlertTypeBlockNotificationReviewRejected, PolicyNewParamsAlertTypeBrandProtectionAlert, PolicyNewParamsAlertTypeBrandProtectionDigest, PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly, PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly, PolicyNewParamsAlertTypeCustomSSLCertificateEventType, PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType, PolicyNewParamsAlertTypeDosAttackL4, PolicyNewParamsAlertTypeDosAttackL7, PolicyNewParamsAlertTypeExpiringServiceTokenAlert, PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert, PolicyNewParamsAlertTypeFbmAutoAdvertisement, PolicyNewParamsAlertTypeFbmDosdAttack, PolicyNewParamsAlertTypeFbmVolumetricAttack, PolicyNewParamsAlertTypeHealthCheckStatusNotification, PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType, PolicyNewParamsAlertTypeHTTPAlertEdgeError, PolicyNewParamsAlertTypeHTTPAlertOriginError, PolicyNewParamsAlertTypeIncidentAlert, PolicyNewParamsAlertTypeLoadBalancingHealthAlert, PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert, PolicyNewParamsAlertTypeLogoMatchAlert, PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent, PolicyNewParamsAlertTypeMaintenanceEventNotification, PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType, PolicyNewParamsAlertTypePagesEventAlert, PolicyNewParamsAlertTypeRadarNotification, PolicyNewParamsAlertTypeRealOriginMonitoring, PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections, PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, PolicyNewParamsAlertTypeScriptmonitorAlertNewResources, PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing, PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing, PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated, PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning, PolicyNewParamsAlertTypeSentinelAlert, PolicyNewParamsAlertTypeStreamLiveNotifications, PolicyNewParamsAlertTypeTrafficAnomaliesAlert, PolicyNewParamsAlertTypeTunnelHealthEvent, PolicyNewParamsAlertTypeTunnelUpdateEvent, PolicyNewParamsAlertTypeUniversalSSLEventType, PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate, PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

type PolicyNewParamsMechanisms struct {
	// UUID
	ID param.Field[PolicyNewParamsMechanismsIDUnion] `json:"id"`
}

func (r PolicyNewParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type PolicyNewParamsMechanismsIDUnion interface {
	ImplementsAlertingPolicyNewParamsMechanismsIDUnion()
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type PolicyNewParamsFilters struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert. A list of identifiers for each component to
	// monitor.
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue param.Field[[]PolicyNewParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled param.Field[[]string] `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment param.Field[[]string] `json:"environment"`
	// Used for configuring pages_event_alert
	Event param.Field[[]string] `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource param.Field[[]string] `json:"event_source"`
	// Usage depends on specific alert type
	EventType param.Field[[]string] `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy param.Field[[]string] `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID param.Field[[]string] `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact param.Field[[]PolicyNewParamsFiltersIncidentImpact] `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID param.Field[[]string] `json:"input_id"`
	// Used for configuring billing_usage_alert
	Limit param.Field[[]string] `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag param.Field[[]string] `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond param.Field[[]string] `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth param.Field[[]string] `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus param.Field[[]string] `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond param.Field[[]string] `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID param.Field[[]string] `json:"pool_id"`
	// Used for configuring billing_usage_alert
	Product param.Field[[]string] `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID param.Field[[]string] `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol param.Field[[]string] `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag param.Field[[]string] `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond param.Field[[]string] `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors param.Field[[]string] `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services param.Field[[]string] `json:"services"`
	// Usage depends on specific alert type
	Slo param.Field[[]string] `json:"slo"`
	// Used for configuring health_check_status_notification
	Status param.Field[[]string] `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname param.Field[[]string] `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP param.Field[[]string] `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName param.Field[[]string] `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions param.Field[[]PolicyNewParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r PolicyNewParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PolicyNewParamsFiltersAlertTriggerPreferencesValue string

const (
	PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0 PolicyNewParamsFiltersAlertTriggerPreferencesValue = "99.0"
	PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0 PolicyNewParamsFiltersAlertTriggerPreferencesValue = "98.0"
	PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0 PolicyNewParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

func (r PolicyNewParamsFiltersAlertTriggerPreferencesValue) IsKnown() bool {
	switch r {
	case PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0, PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0, PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0:
		return true
	}
	return false
}

type PolicyNewParamsFiltersIncidentImpact string

const (
	PolicyNewParamsFiltersIncidentImpactIncidentImpactNone     PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor    PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor    PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	PolicyNewParamsFiltersIncidentImpactIncidentImpactCritical PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

func (r PolicyNewParamsFiltersIncidentImpact) IsKnown() bool {
	switch r {
	case PolicyNewParamsFiltersIncidentImpactIncidentImpactNone, PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor, PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor, PolicyNewParamsFiltersIncidentImpactIncidentImpactCritical:
		return true
	}
	return false
}

type PolicyNewParamsFiltersTrafficExclusion string

const (
	PolicyNewParamsFiltersTrafficExclusionSecurityEvents PolicyNewParamsFiltersTrafficExclusion = "security_events"
)

func (r PolicyNewParamsFiltersTrafficExclusion) IsKnown() bool {
	switch r {
	case PolicyNewParamsFiltersTrafficExclusionSecurityEvents:
		return true
	}
	return false
}

type PolicyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PolicyNewResponse     `json:"result,required"`
	// Whether the API call was successful
	Success PolicyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    policyNewResponseEnvelopeJSON    `json:"-"`
}

// policyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyNewResponseEnvelope]
type policyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyNewResponseEnvelopeSuccess bool

const (
	PolicyNewResponseEnvelopeSuccessTrue PolicyNewResponseEnvelopeSuccess = true
)

func (r PolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[PolicyUpdateParamsAlertType] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[PolicyUpdateParamsFilters] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]PolicyUpdateParamsMechanisms] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type PolicyUpdateParamsAlertType string

const (
	PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType         PolicyUpdateParamsAlertType = "access_custom_certificate_expiration_type"
	PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL4Alert                     PolicyUpdateParamsAlertType = "advanced_ddos_attack_l4_alert"
	PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL7Alert                     PolicyUpdateParamsAlertType = "advanced_ddos_attack_l7_alert"
	PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError                        PolicyUpdateParamsAlertType = "advanced_http_alert_error"
	PolicyUpdateParamsAlertTypeBGPHijackNotification                         PolicyUpdateParamsAlertType = "bgp_hijack_notification"
	PolicyUpdateParamsAlertTypeBillingUsageAlert                             PolicyUpdateParamsAlertType = "billing_usage_alert"
	PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved                 PolicyUpdateParamsAlertType = "block_notification_block_removed"
	PolicyUpdateParamsAlertTypeBlockNotificationNewBlock                     PolicyUpdateParamsAlertType = "block_notification_new_block"
	PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected               PolicyUpdateParamsAlertType = "block_notification_review_rejected"
	PolicyUpdateParamsAlertTypeBrandProtectionAlert                          PolicyUpdateParamsAlertType = "brand_protection_alert"
	PolicyUpdateParamsAlertTypeBrandProtectionDigest                         PolicyUpdateParamsAlertType = "brand_protection_digest"
	PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly                      PolicyUpdateParamsAlertType = "clickhouse_alert_fw_anomaly"
	PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly                   PolicyUpdateParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType                 PolicyUpdateParamsAlertType = "custom_ssl_certificate_event_type"
	PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType              PolicyUpdateParamsAlertType = "dedicated_ssl_certificate_event_type"
	PolicyUpdateParamsAlertTypeDosAttackL4                                   PolicyUpdateParamsAlertType = "dos_attack_l4"
	PolicyUpdateParamsAlertTypeDosAttackL7                                   PolicyUpdateParamsAlertType = "dos_attack_l7"
	PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert                     PolicyUpdateParamsAlertType = "expiring_service_token_alert"
	PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert                PolicyUpdateParamsAlertType = "failing_logpush_job_disabled_alert"
	PolicyUpdateParamsAlertTypeFbmAutoAdvertisement                          PolicyUpdateParamsAlertType = "fbm_auto_advertisement"
	PolicyUpdateParamsAlertTypeFbmDosdAttack                                 PolicyUpdateParamsAlertType = "fbm_dosd_attack"
	PolicyUpdateParamsAlertTypeFbmVolumetricAttack                           PolicyUpdateParamsAlertType = "fbm_volumetric_attack"
	PolicyUpdateParamsAlertTypeHealthCheckStatusNotification                 PolicyUpdateParamsAlertType = "health_check_status_notification"
	PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType    PolicyUpdateParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	PolicyUpdateParamsAlertTypeHTTPAlertEdgeError                            PolicyUpdateParamsAlertType = "http_alert_edge_error"
	PolicyUpdateParamsAlertTypeHTTPAlertOriginError                          PolicyUpdateParamsAlertType = "http_alert_origin_error"
	PolicyUpdateParamsAlertTypeIncidentAlert                                 PolicyUpdateParamsAlertType = "incident_alert"
	PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert                      PolicyUpdateParamsAlertType = "load_balancing_health_alert"
	PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert              PolicyUpdateParamsAlertType = "load_balancing_pool_enablement_alert"
	PolicyUpdateParamsAlertTypeLogoMatchAlert                                PolicyUpdateParamsAlertType = "logo_match_alert"
	PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent                   PolicyUpdateParamsAlertType = "magic_tunnel_health_check_event"
	PolicyUpdateParamsAlertTypeMaintenanceEventNotification                  PolicyUpdateParamsAlertType = "maintenance_event_notification"
	PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType PolicyUpdateParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	PolicyUpdateParamsAlertTypePagesEventAlert                               PolicyUpdateParamsAlertType = "pages_event_alert"
	PolicyUpdateParamsAlertTypeRadarNotification                             PolicyUpdateParamsAlertType = "radar_notification"
	PolicyUpdateParamsAlertTypeRealOriginMonitoring                          PolicyUpdateParamsAlertType = "real_origin_monitoring"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts                    PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_hosts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL             PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources                PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_resources"
	PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing               PolicyUpdateParamsAlertType = "secondary_dns_all_primaries_failing"
	PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing                  PolicyUpdateParamsAlertType = "secondary_dns_primaries_failing"
	PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           PolicyUpdateParamsAlertType = "secondary_dns_zone_successfully_updated"
	PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning             PolicyUpdateParamsAlertType = "secondary_dns_zone_validation_warning"
	PolicyUpdateParamsAlertTypeSentinelAlert                                 PolicyUpdateParamsAlertType = "sentinel_alert"
	PolicyUpdateParamsAlertTypeStreamLiveNotifications                       PolicyUpdateParamsAlertType = "stream_live_notifications"
	PolicyUpdateParamsAlertTypeTrafficAnomaliesAlert                         PolicyUpdateParamsAlertType = "traffic_anomalies_alert"
	PolicyUpdateParamsAlertTypeTunnelHealthEvent                             PolicyUpdateParamsAlertType = "tunnel_health_event"
	PolicyUpdateParamsAlertTypeTunnelUpdateEvent                             PolicyUpdateParamsAlertType = "tunnel_update_event"
	PolicyUpdateParamsAlertTypeUniversalSSLEventType                         PolicyUpdateParamsAlertType = "universal_ssl_event_type"
	PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate                     PolicyUpdateParamsAlertType = "web_analytics_metrics_update"
	PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType        PolicyUpdateParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r PolicyUpdateParamsAlertType) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType, PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL4Alert, PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL7Alert, PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError, PolicyUpdateParamsAlertTypeBGPHijackNotification, PolicyUpdateParamsAlertTypeBillingUsageAlert, PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved, PolicyUpdateParamsAlertTypeBlockNotificationNewBlock, PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected, PolicyUpdateParamsAlertTypeBrandProtectionAlert, PolicyUpdateParamsAlertTypeBrandProtectionDigest, PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly, PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly, PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType, PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType, PolicyUpdateParamsAlertTypeDosAttackL4, PolicyUpdateParamsAlertTypeDosAttackL7, PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert, PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert, PolicyUpdateParamsAlertTypeFbmAutoAdvertisement, PolicyUpdateParamsAlertTypeFbmDosdAttack, PolicyUpdateParamsAlertTypeFbmVolumetricAttack, PolicyUpdateParamsAlertTypeHealthCheckStatusNotification, PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType, PolicyUpdateParamsAlertTypeHTTPAlertEdgeError, PolicyUpdateParamsAlertTypeHTTPAlertOriginError, PolicyUpdateParamsAlertTypeIncidentAlert, PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert, PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert, PolicyUpdateParamsAlertTypeLogoMatchAlert, PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent, PolicyUpdateParamsAlertTypeMaintenanceEventNotification, PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType, PolicyUpdateParamsAlertTypePagesEventAlert, PolicyUpdateParamsAlertTypeRadarNotification, PolicyUpdateParamsAlertTypeRealOriginMonitoring, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources, PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing, PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing, PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated, PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning, PolicyUpdateParamsAlertTypeSentinelAlert, PolicyUpdateParamsAlertTypeStreamLiveNotifications, PolicyUpdateParamsAlertTypeTrafficAnomaliesAlert, PolicyUpdateParamsAlertTypeTunnelHealthEvent, PolicyUpdateParamsAlertTypeTunnelUpdateEvent, PolicyUpdateParamsAlertTypeUniversalSSLEventType, PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate, PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type PolicyUpdateParamsFilters struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert. A list of identifiers for each component to
	// monitor.
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue param.Field[[]PolicyUpdateParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled param.Field[[]string] `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment param.Field[[]string] `json:"environment"`
	// Used for configuring pages_event_alert
	Event param.Field[[]string] `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource param.Field[[]string] `json:"event_source"`
	// Usage depends on specific alert type
	EventType param.Field[[]string] `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy param.Field[[]string] `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID param.Field[[]string] `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact param.Field[[]PolicyUpdateParamsFiltersIncidentImpact] `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID param.Field[[]string] `json:"input_id"`
	// Used for configuring billing_usage_alert
	Limit param.Field[[]string] `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag param.Field[[]string] `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond param.Field[[]string] `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth param.Field[[]string] `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus param.Field[[]string] `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond param.Field[[]string] `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID param.Field[[]string] `json:"pool_id"`
	// Used for configuring billing_usage_alert
	Product param.Field[[]string] `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID param.Field[[]string] `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol param.Field[[]string] `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag param.Field[[]string] `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond param.Field[[]string] `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors param.Field[[]string] `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services param.Field[[]string] `json:"services"`
	// Usage depends on specific alert type
	Slo param.Field[[]string] `json:"slo"`
	// Used for configuring health_check_status_notification
	Status param.Field[[]string] `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname param.Field[[]string] `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP param.Field[[]string] `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName param.Field[[]string] `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions param.Field[[]PolicyUpdateParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r PolicyUpdateParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PolicyUpdateParamsFiltersAlertTriggerPreferencesValue string

const (
	PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0 PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "99.0"
	PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0 PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "98.0"
	PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0 PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

func (r PolicyUpdateParamsFiltersAlertTriggerPreferencesValue) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0, PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0, PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0:
		return true
	}
	return false
}

type PolicyUpdateParamsFiltersIncidentImpact string

const (
	PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone     PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor    PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor    PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	PolicyUpdateParamsFiltersIncidentImpactIncidentImpactCritical PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

func (r PolicyUpdateParamsFiltersIncidentImpact) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone, PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor, PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor, PolicyUpdateParamsFiltersIncidentImpactIncidentImpactCritical:
		return true
	}
	return false
}

type PolicyUpdateParamsFiltersTrafficExclusion string

const (
	PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents PolicyUpdateParamsFiltersTrafficExclusion = "security_events"
)

func (r PolicyUpdateParamsFiltersTrafficExclusion) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents:
		return true
	}
	return false
}

type PolicyUpdateParamsMechanisms struct {
	// UUID
	ID param.Field[PolicyUpdateParamsMechanismsIDUnion] `json:"id"`
}

func (r PolicyUpdateParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type PolicyUpdateParamsMechanismsIDUnion interface {
	ImplementsAlertingPolicyUpdateParamsMechanismsIDUnion()
}

type PolicyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PolicyUpdateResponse  `json:"result,required"`
	// Whether the API call was successful
	Success PolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    policyUpdateResponseEnvelopeJSON    `json:"-"`
}

// policyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyUpdateResponseEnvelope]
type policyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyUpdateResponseEnvelopeSuccess bool

const (
	PolicyUpdateResponseEnvelopeSuccessTrue PolicyUpdateResponseEnvelopeSuccess = true
)

func (r PolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   PolicyDeleteResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PolicyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PolicyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       policyDeleteResponseEnvelopeJSON       `json:"-"`
}

// policyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyDeleteResponseEnvelope]
type policyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyDeleteResponseEnvelopeSuccess bool

const (
	PolicyDeleteResponseEnvelopeSuccessTrue PolicyDeleteResponseEnvelopeSuccess = true
)

func (r PolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       policyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// policyDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [PolicyDeleteResponseEnvelopeResultInfo]
type policyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PolicyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AlertingPolicies      `json:"result,required"`
	// Whether the API call was successful
	Success PolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    policyGetResponseEnvelopeJSON    `json:"-"`
}

// policyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyGetResponseEnvelope]
type policyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyGetResponseEnvelopeSuccess bool

const (
	PolicyGetResponseEnvelopeSuccessTrue PolicyGetResponseEnvelopeSuccess = true
)

func (r PolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

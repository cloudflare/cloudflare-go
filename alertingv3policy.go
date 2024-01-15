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

// AlertingV3PolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAlertingV3PolicyService] method
// instead.
type AlertingV3PolicyService struct {
	Options []option.RequestOption
}

// NewAlertingV3PolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlertingV3PolicyService(opts ...option.RequestOption) (r *AlertingV3PolicyService) {
	r = &AlertingV3PolicyService{}
	r.Options = opts
	return
}

// Creates a new Notification policy.
func (r *AlertingV3PolicyService) New(ctx context.Context, accountID string, body AlertingV3PolicyNewParams, opts ...option.RequestOption) (res *AlertingV3PolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details for a single policy.
func (r *AlertingV3PolicyService) Get(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AlertingV3PolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Notification policy.
func (r *AlertingV3PolicyService) Update(ctx context.Context, accountID string, policyID string, body AlertingV3PolicyUpdateParams, opts ...option.RequestOption) (res *AlertingV3PolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a list of all Notification policies.
func (r *AlertingV3PolicyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AlertingV3PolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Notification policy.
func (r *AlertingV3PolicyService) Delete(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AlertingV3PolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AlertingV3PolicyNewResponse struct {
	Errors   []AlertingV3PolicyNewResponseError   `json:"errors"`
	Messages []AlertingV3PolicyNewResponseMessage `json:"messages"`
	Result   AlertingV3PolicyNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AlertingV3PolicyNewResponseSuccess `json:"success"`
	JSON    alertingV3PolicyNewResponseJSON    `json:"-"`
}

// alertingV3PolicyNewResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyNewResponse]
type alertingV3PolicyNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNewResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    alertingV3PolicyNewResponseErrorJSON `json:"-"`
}

// alertingV3PolicyNewResponseErrorJSON contains the JSON metadata for the struct
// [AlertingV3PolicyNewResponseError]
type alertingV3PolicyNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNewResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    alertingV3PolicyNewResponseMessageJSON `json:"-"`
}

// alertingV3PolicyNewResponseMessageJSON contains the JSON metadata for the struct
// [AlertingV3PolicyNewResponseMessage]
type alertingV3PolicyNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNewResponseResult struct {
	// UUID
	ID   string                                `json:"id"`
	JSON alertingV3PolicyNewResponseResultJSON `json:"-"`
}

// alertingV3PolicyNewResponseResultJSON contains the JSON metadata for the struct
// [AlertingV3PolicyNewResponseResult]
type alertingV3PolicyNewResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyNewResponseSuccess bool

const (
	AlertingV3PolicyNewResponseSuccessTrue AlertingV3PolicyNewResponseSuccess = true
)

type AlertingV3PolicyGetResponse struct {
	Errors   []AlertingV3PolicyGetResponseError   `json:"errors"`
	Messages []AlertingV3PolicyGetResponseMessage `json:"messages"`
	Result   AlertingV3PolicyGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AlertingV3PolicyGetResponseSuccess `json:"success"`
	JSON    alertingV3PolicyGetResponseJSON    `json:"-"`
}

// alertingV3PolicyGetResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponse]
type alertingV3PolicyGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    alertingV3PolicyGetResponseErrorJSON `json:"-"`
}

// alertingV3PolicyGetResponseErrorJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponseError]
type alertingV3PolicyGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    alertingV3PolicyGetResponseMessageJSON `json:"-"`
}

// alertingV3PolicyGetResponseMessageJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponseMessage]
type alertingV3PolicyGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseResult struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertingV3PolicyGetResponseResultAlertType `json:"alert_type"`
	Created   time.Time                                  `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertingV3PolicyGetResponseResultFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms interface{} `json:"mechanisms"`
	Modified   time.Time   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                                `json:"name"`
	JSON alertingV3PolicyGetResponseResultJSON `json:"-"`
}

// alertingV3PolicyGetResponseResultJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponseResult]
type alertingV3PolicyGetResponseResultJSON struct {
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

func (r *AlertingV3PolicyGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyGetResponseResultAlertType string

const (
	AlertingV3PolicyGetResponseResultAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyGetResponseResultAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyGetResponseResultAlertTypeAdvancedDdosAttackL4Alert                     AlertingV3PolicyGetResponseResultAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyGetResponseResultAlertTypeAdvancedDdosAttackL7Alert                     AlertingV3PolicyGetResponseResultAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyGetResponseResultAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyGetResponseResultAlertType = "advanced_http_alert_error"
	AlertingV3PolicyGetResponseResultAlertTypeBgpHijackNotification                         AlertingV3PolicyGetResponseResultAlertType = "bgp_hijack_notification"
	AlertingV3PolicyGetResponseResultAlertTypeBillingUsageAlert                             AlertingV3PolicyGetResponseResultAlertType = "billing_usage_alert"
	AlertingV3PolicyGetResponseResultAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyGetResponseResultAlertType = "block_notification_block_removed"
	AlertingV3PolicyGetResponseResultAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyGetResponseResultAlertType = "block_notification_new_block"
	AlertingV3PolicyGetResponseResultAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyGetResponseResultAlertType = "block_notification_review_rejected"
	AlertingV3PolicyGetResponseResultAlertTypeBrandProtectionAlert                          AlertingV3PolicyGetResponseResultAlertType = "brand_protection_alert"
	AlertingV3PolicyGetResponseResultAlertTypeBrandProtectionDigest                         AlertingV3PolicyGetResponseResultAlertType = "brand_protection_digest"
	AlertingV3PolicyGetResponseResultAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyGetResponseResultAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyGetResponseResultAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyGetResponseResultAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyGetResponseResultAlertTypeCustomSslCertificateEventType                 AlertingV3PolicyGetResponseResultAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyGetResponseResultAlertTypeDedicatedSslCertificateEventType              AlertingV3PolicyGetResponseResultAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyGetResponseResultAlertTypeDosAttackL4                                   AlertingV3PolicyGetResponseResultAlertType = "dos_attack_l4"
	AlertingV3PolicyGetResponseResultAlertTypeDosAttackL7                                   AlertingV3PolicyGetResponseResultAlertType = "dos_attack_l7"
	AlertingV3PolicyGetResponseResultAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyGetResponseResultAlertType = "expiring_service_token_alert"
	AlertingV3PolicyGetResponseResultAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyGetResponseResultAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyGetResponseResultAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyGetResponseResultAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyGetResponseResultAlertTypeFbmDosdAttack                                 AlertingV3PolicyGetResponseResultAlertType = "fbm_dosd_attack"
	AlertingV3PolicyGetResponseResultAlertTypeFbmVolumetricAttack                           AlertingV3PolicyGetResponseResultAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyGetResponseResultAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyGetResponseResultAlertType = "health_check_status_notification"
	AlertingV3PolicyGetResponseResultAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyGetResponseResultAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyGetResponseResultAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyGetResponseResultAlertType = "http_alert_edge_error"
	AlertingV3PolicyGetResponseResultAlertTypeHTTPAlertOriginError                          AlertingV3PolicyGetResponseResultAlertType = "http_alert_origin_error"
	AlertingV3PolicyGetResponseResultAlertTypeIncidentAlert                                 AlertingV3PolicyGetResponseResultAlertType = "incident_alert"
	AlertingV3PolicyGetResponseResultAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyGetResponseResultAlertType = "load_balancing_health_alert"
	AlertingV3PolicyGetResponseResultAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyGetResponseResultAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyGetResponseResultAlertTypeLogoMatchAlert                                AlertingV3PolicyGetResponseResultAlertType = "logo_match_alert"
	AlertingV3PolicyGetResponseResultAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyGetResponseResultAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyGetResponseResultAlertTypeMaintenanceEventNotification                  AlertingV3PolicyGetResponseResultAlertType = "maintenance_event_notification"
	AlertingV3PolicyGetResponseResultAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyGetResponseResultAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyGetResponseResultAlertTypePagesEventAlert                               AlertingV3PolicyGetResponseResultAlertType = "pages_event_alert"
	AlertingV3PolicyGetResponseResultAlertTypeRadarNotification                             AlertingV3PolicyGetResponseResultAlertType = "radar_notification"
	AlertingV3PolicyGetResponseResultAlertTypeRealOriginMonitoring                          AlertingV3PolicyGetResponseResultAlertType = "real_origin_monitoring"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyGetResponseResultAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyGetResponseResultAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyGetResponseResultAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyGetResponseResultAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyGetResponseResultAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyGetResponseResultAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyGetResponseResultAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyGetResponseResultAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyGetResponseResultAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyGetResponseResultAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyGetResponseResultAlertTypeSentinelAlert                                 AlertingV3PolicyGetResponseResultAlertType = "sentinel_alert"
	AlertingV3PolicyGetResponseResultAlertTypeStreamLiveNotifications                       AlertingV3PolicyGetResponseResultAlertType = "stream_live_notifications"
	AlertingV3PolicyGetResponseResultAlertTypeTunnelHealthEvent                             AlertingV3PolicyGetResponseResultAlertType = "tunnel_health_event"
	AlertingV3PolicyGetResponseResultAlertTypeTunnelUpdateEvent                             AlertingV3PolicyGetResponseResultAlertType = "tunnel_update_event"
	AlertingV3PolicyGetResponseResultAlertTypeUniversalSslEventType                         AlertingV3PolicyGetResponseResultAlertType = "universal_ssl_event_type"
	AlertingV3PolicyGetResponseResultAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyGetResponseResultAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyGetResponseResultAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyGetResponseResultAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyGetResponseResultFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue []AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []AlertingV3PolicyGetResponseResultFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []AlertingV3PolicyGetResponseResultFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                                     `json:"zones"`
	JSON  alertingV3PolicyGetResponseResultFiltersJSON `json:"-"`
}

// alertingV3PolicyGetResponseResultFiltersJSON contains the JSON metadata for the
// struct [AlertingV3PolicyGetResponseResultFilters]
type alertingV3PolicyGetResponseResultFiltersJSON struct {
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

func (r *AlertingV3PolicyGetResponseResultFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyGetResponseResultFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyGetResponseResultFiltersIncidentImpact string

const (
	AlertingV3PolicyGetResponseResultFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyGetResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyGetResponseResultFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyGetResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyGetResponseResultFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyGetResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyGetResponseResultFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyGetResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyGetResponseResultFiltersTrafficExclusion string

const (
	AlertingV3PolicyGetResponseResultFiltersTrafficExclusionSecurityEvents AlertingV3PolicyGetResponseResultFiltersTrafficExclusion = "security_events"
)

// Whether the API call was successful
type AlertingV3PolicyGetResponseSuccess bool

const (
	AlertingV3PolicyGetResponseSuccessTrue AlertingV3PolicyGetResponseSuccess = true
)

type AlertingV3PolicyUpdateResponse struct {
	Errors   []AlertingV3PolicyUpdateResponseError   `json:"errors"`
	Messages []AlertingV3PolicyUpdateResponseMessage `json:"messages"`
	Result   AlertingV3PolicyUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AlertingV3PolicyUpdateResponseSuccess `json:"success"`
	JSON    alertingV3PolicyUpdateResponseJSON    `json:"-"`
}

// alertingV3PolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyUpdateResponse]
type alertingV3PolicyUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    alertingV3PolicyUpdateResponseErrorJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AlertingV3PolicyUpdateResponseError]
type alertingV3PolicyUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    alertingV3PolicyUpdateResponseMessageJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AlertingV3PolicyUpdateResponseMessage]
type alertingV3PolicyUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyUpdateResponseResult struct {
	// UUID
	ID   string                                   `json:"id"`
	JSON alertingV3PolicyUpdateResponseResultJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseResultJSON contains the JSON metadata for the
// struct [AlertingV3PolicyUpdateResponseResult]
type alertingV3PolicyUpdateResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyUpdateResponseSuccess bool

const (
	AlertingV3PolicyUpdateResponseSuccessTrue AlertingV3PolicyUpdateResponseSuccess = true
)

type AlertingV3PolicyListResponse struct {
	Errors     []AlertingV3PolicyListResponseError    `json:"errors"`
	Messages   []AlertingV3PolicyListResponseMessage  `json:"messages"`
	Result     []AlertingV3PolicyListResponseResult   `json:"result"`
	ResultInfo AlertingV3PolicyListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AlertingV3PolicyListResponseSuccess `json:"success"`
	JSON    alertingV3PolicyListResponseJSON    `json:"-"`
}

// alertingV3PolicyListResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyListResponse]
type alertingV3PolicyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    alertingV3PolicyListResponseErrorJSON `json:"-"`
}

// alertingV3PolicyListResponseErrorJSON contains the JSON metadata for the struct
// [AlertingV3PolicyListResponseError]
type alertingV3PolicyListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    alertingV3PolicyListResponseMessageJSON `json:"-"`
}

// alertingV3PolicyListResponseMessageJSON contains the JSON metadata for the
// struct [AlertingV3PolicyListResponseMessage]
type alertingV3PolicyListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseResult struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertingV3PolicyListResponseResultAlertType `json:"alert_type"`
	Created   time.Time                                   `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertingV3PolicyListResponseResultFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms interface{} `json:"mechanisms"`
	Modified   time.Time   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                                 `json:"name"`
	JSON alertingV3PolicyListResponseResultJSON `json:"-"`
}

// alertingV3PolicyListResponseResultJSON contains the JSON metadata for the struct
// [AlertingV3PolicyListResponseResult]
type alertingV3PolicyListResponseResultJSON struct {
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

func (r *AlertingV3PolicyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyListResponseResultAlertType string

const (
	AlertingV3PolicyListResponseResultAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyListResponseResultAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyListResponseResultAlertTypeAdvancedDdosAttackL4Alert                     AlertingV3PolicyListResponseResultAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyListResponseResultAlertTypeAdvancedDdosAttackL7Alert                     AlertingV3PolicyListResponseResultAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyListResponseResultAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyListResponseResultAlertType = "advanced_http_alert_error"
	AlertingV3PolicyListResponseResultAlertTypeBgpHijackNotification                         AlertingV3PolicyListResponseResultAlertType = "bgp_hijack_notification"
	AlertingV3PolicyListResponseResultAlertTypeBillingUsageAlert                             AlertingV3PolicyListResponseResultAlertType = "billing_usage_alert"
	AlertingV3PolicyListResponseResultAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyListResponseResultAlertType = "block_notification_block_removed"
	AlertingV3PolicyListResponseResultAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyListResponseResultAlertType = "block_notification_new_block"
	AlertingV3PolicyListResponseResultAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyListResponseResultAlertType = "block_notification_review_rejected"
	AlertingV3PolicyListResponseResultAlertTypeBrandProtectionAlert                          AlertingV3PolicyListResponseResultAlertType = "brand_protection_alert"
	AlertingV3PolicyListResponseResultAlertTypeBrandProtectionDigest                         AlertingV3PolicyListResponseResultAlertType = "brand_protection_digest"
	AlertingV3PolicyListResponseResultAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyListResponseResultAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyListResponseResultAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyListResponseResultAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyListResponseResultAlertTypeCustomSslCertificateEventType                 AlertingV3PolicyListResponseResultAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyListResponseResultAlertTypeDedicatedSslCertificateEventType              AlertingV3PolicyListResponseResultAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyListResponseResultAlertTypeDosAttackL4                                   AlertingV3PolicyListResponseResultAlertType = "dos_attack_l4"
	AlertingV3PolicyListResponseResultAlertTypeDosAttackL7                                   AlertingV3PolicyListResponseResultAlertType = "dos_attack_l7"
	AlertingV3PolicyListResponseResultAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyListResponseResultAlertType = "expiring_service_token_alert"
	AlertingV3PolicyListResponseResultAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyListResponseResultAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyListResponseResultAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyListResponseResultAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyListResponseResultAlertTypeFbmDosdAttack                                 AlertingV3PolicyListResponseResultAlertType = "fbm_dosd_attack"
	AlertingV3PolicyListResponseResultAlertTypeFbmVolumetricAttack                           AlertingV3PolicyListResponseResultAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyListResponseResultAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyListResponseResultAlertType = "health_check_status_notification"
	AlertingV3PolicyListResponseResultAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyListResponseResultAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyListResponseResultAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyListResponseResultAlertType = "http_alert_edge_error"
	AlertingV3PolicyListResponseResultAlertTypeHTTPAlertOriginError                          AlertingV3PolicyListResponseResultAlertType = "http_alert_origin_error"
	AlertingV3PolicyListResponseResultAlertTypeIncidentAlert                                 AlertingV3PolicyListResponseResultAlertType = "incident_alert"
	AlertingV3PolicyListResponseResultAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyListResponseResultAlertType = "load_balancing_health_alert"
	AlertingV3PolicyListResponseResultAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyListResponseResultAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyListResponseResultAlertTypeLogoMatchAlert                                AlertingV3PolicyListResponseResultAlertType = "logo_match_alert"
	AlertingV3PolicyListResponseResultAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyListResponseResultAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyListResponseResultAlertTypeMaintenanceEventNotification                  AlertingV3PolicyListResponseResultAlertType = "maintenance_event_notification"
	AlertingV3PolicyListResponseResultAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyListResponseResultAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyListResponseResultAlertTypePagesEventAlert                               AlertingV3PolicyListResponseResultAlertType = "pages_event_alert"
	AlertingV3PolicyListResponseResultAlertTypeRadarNotification                             AlertingV3PolicyListResponseResultAlertType = "radar_notification"
	AlertingV3PolicyListResponseResultAlertTypeRealOriginMonitoring                          AlertingV3PolicyListResponseResultAlertType = "real_origin_monitoring"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyListResponseResultAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyListResponseResultAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyListResponseResultAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyListResponseResultAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyListResponseResultAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyListResponseResultAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyListResponseResultAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyListResponseResultAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyListResponseResultAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyListResponseResultAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyListResponseResultAlertTypeSentinelAlert                                 AlertingV3PolicyListResponseResultAlertType = "sentinel_alert"
	AlertingV3PolicyListResponseResultAlertTypeStreamLiveNotifications                       AlertingV3PolicyListResponseResultAlertType = "stream_live_notifications"
	AlertingV3PolicyListResponseResultAlertTypeTunnelHealthEvent                             AlertingV3PolicyListResponseResultAlertType = "tunnel_health_event"
	AlertingV3PolicyListResponseResultAlertTypeTunnelUpdateEvent                             AlertingV3PolicyListResponseResultAlertType = "tunnel_update_event"
	AlertingV3PolicyListResponseResultAlertTypeUniversalSslEventType                         AlertingV3PolicyListResponseResultAlertType = "universal_ssl_event_type"
	AlertingV3PolicyListResponseResultAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyListResponseResultAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyListResponseResultAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyListResponseResultAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyListResponseResultFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue []AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []AlertingV3PolicyListResponseResultFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []AlertingV3PolicyListResponseResultFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                                      `json:"zones"`
	JSON  alertingV3PolicyListResponseResultFiltersJSON `json:"-"`
}

// alertingV3PolicyListResponseResultFiltersJSON contains the JSON metadata for the
// struct [AlertingV3PolicyListResponseResultFilters]
type alertingV3PolicyListResponseResultFiltersJSON struct {
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

func (r *AlertingV3PolicyListResponseResultFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyListResponseResultFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyListResponseResultFiltersIncidentImpact string

const (
	AlertingV3PolicyListResponseResultFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyListResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyListResponseResultFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyListResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyListResponseResultFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyListResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyListResponseResultFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyListResponseResultFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyListResponseResultFiltersTrafficExclusion string

const (
	AlertingV3PolicyListResponseResultFiltersTrafficExclusionSecurityEvents AlertingV3PolicyListResponseResultFiltersTrafficExclusion = "security_events"
)

type AlertingV3PolicyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       alertingV3PolicyListResponseResultInfoJSON `json:"-"`
}

// alertingV3PolicyListResponseResultInfoJSON contains the JSON metadata for the
// struct [AlertingV3PolicyListResponseResultInfo]
type alertingV3PolicyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyListResponseSuccess bool

const (
	AlertingV3PolicyListResponseSuccessTrue AlertingV3PolicyListResponseSuccess = true
)

type AlertingV3PolicyDeleteResponse struct {
	Errors     []AlertingV3PolicyDeleteResponseError    `json:"errors"`
	Messages   []AlertingV3PolicyDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                            `json:"result,nullable"`
	ResultInfo AlertingV3PolicyDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AlertingV3PolicyDeleteResponseSuccess `json:"success"`
	JSON    alertingV3PolicyDeleteResponseJSON    `json:"-"`
}

// alertingV3PolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyDeleteResponse]
type alertingV3PolicyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    alertingV3PolicyDeleteResponseErrorJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AlertingV3PolicyDeleteResponseError]
type alertingV3PolicyDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    alertingV3PolicyDeleteResponseMessageJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AlertingV3PolicyDeleteResponseMessage]
type alertingV3PolicyDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       alertingV3PolicyDeleteResponseResultInfoJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AlertingV3PolicyDeleteResponseResultInfo]
type alertingV3PolicyDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyDeleteResponseSuccess bool

const (
	AlertingV3PolicyDeleteResponseSuccessTrue AlertingV3PolicyDeleteResponseSuccess = true
)

type AlertingV3PolicyNewParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertingV3PolicyNewParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[interface{}] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[AlertingV3PolicyNewParamsFilters] `json:"filters"`
}

func (r AlertingV3PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyNewParamsAlertType string

const (
	AlertingV3PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyNewParamsAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyNewParamsAlertTypeAdvancedDdosAttackL4Alert                     AlertingV3PolicyNewParamsAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyNewParamsAlertTypeAdvancedDdosAttackL7Alert                     AlertingV3PolicyNewParamsAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyNewParamsAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyNewParamsAlertType = "advanced_http_alert_error"
	AlertingV3PolicyNewParamsAlertTypeBgpHijackNotification                         AlertingV3PolicyNewParamsAlertType = "bgp_hijack_notification"
	AlertingV3PolicyNewParamsAlertTypeBillingUsageAlert                             AlertingV3PolicyNewParamsAlertType = "billing_usage_alert"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyNewParamsAlertType = "block_notification_block_removed"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyNewParamsAlertType = "block_notification_new_block"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyNewParamsAlertType = "block_notification_review_rejected"
	AlertingV3PolicyNewParamsAlertTypeBrandProtectionAlert                          AlertingV3PolicyNewParamsAlertType = "brand_protection_alert"
	AlertingV3PolicyNewParamsAlertTypeBrandProtectionDigest                         AlertingV3PolicyNewParamsAlertType = "brand_protection_digest"
	AlertingV3PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyNewParamsAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyNewParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyNewParamsAlertTypeCustomSslCertificateEventType                 AlertingV3PolicyNewParamsAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyNewParamsAlertTypeDedicatedSslCertificateEventType              AlertingV3PolicyNewParamsAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyNewParamsAlertTypeDosAttackL4                                   AlertingV3PolicyNewParamsAlertType = "dos_attack_l4"
	AlertingV3PolicyNewParamsAlertTypeDosAttackL7                                   AlertingV3PolicyNewParamsAlertType = "dos_attack_l7"
	AlertingV3PolicyNewParamsAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyNewParamsAlertType = "expiring_service_token_alert"
	AlertingV3PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyNewParamsAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyNewParamsAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyNewParamsAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyNewParamsAlertTypeFbmDosdAttack                                 AlertingV3PolicyNewParamsAlertType = "fbm_dosd_attack"
	AlertingV3PolicyNewParamsAlertTypeFbmVolumetricAttack                           AlertingV3PolicyNewParamsAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyNewParamsAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyNewParamsAlertType = "health_check_status_notification"
	AlertingV3PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyNewParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyNewParamsAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyNewParamsAlertType = "http_alert_edge_error"
	AlertingV3PolicyNewParamsAlertTypeHTTPAlertOriginError                          AlertingV3PolicyNewParamsAlertType = "http_alert_origin_error"
	AlertingV3PolicyNewParamsAlertTypeIncidentAlert                                 AlertingV3PolicyNewParamsAlertType = "incident_alert"
	AlertingV3PolicyNewParamsAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyNewParamsAlertType = "load_balancing_health_alert"
	AlertingV3PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyNewParamsAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyNewParamsAlertTypeLogoMatchAlert                                AlertingV3PolicyNewParamsAlertType = "logo_match_alert"
	AlertingV3PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyNewParamsAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyNewParamsAlertTypeMaintenanceEventNotification                  AlertingV3PolicyNewParamsAlertType = "maintenance_event_notification"
	AlertingV3PolicyNewParamsAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyNewParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyNewParamsAlertTypePagesEventAlert                               AlertingV3PolicyNewParamsAlertType = "pages_event_alert"
	AlertingV3PolicyNewParamsAlertTypeRadarNotification                             AlertingV3PolicyNewParamsAlertType = "radar_notification"
	AlertingV3PolicyNewParamsAlertTypeRealOriginMonitoring                          AlertingV3PolicyNewParamsAlertType = "real_origin_monitoring"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyNewParamsAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyNewParamsAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyNewParamsAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyNewParamsAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyNewParamsAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyNewParamsAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyNewParamsAlertTypeSentinelAlert                                 AlertingV3PolicyNewParamsAlertType = "sentinel_alert"
	AlertingV3PolicyNewParamsAlertTypeStreamLiveNotifications                       AlertingV3PolicyNewParamsAlertType = "stream_live_notifications"
	AlertingV3PolicyNewParamsAlertTypeTunnelHealthEvent                             AlertingV3PolicyNewParamsAlertType = "tunnel_health_event"
	AlertingV3PolicyNewParamsAlertTypeTunnelUpdateEvent                             AlertingV3PolicyNewParamsAlertType = "tunnel_update_event"
	AlertingV3PolicyNewParamsAlertTypeUniversalSslEventType                         AlertingV3PolicyNewParamsAlertType = "universal_ssl_event_type"
	AlertingV3PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyNewParamsAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyNewParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyNewParamsFilters struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue param.Field[[]AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
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
	IncidentImpact param.Field[[]AlertingV3PolicyNewParamsFiltersIncidentImpact] `json:"incident_impact"`
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
	TrafficExclusions param.Field[[]AlertingV3PolicyNewParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r AlertingV3PolicyNewParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyNewParamsFiltersIncidentImpact string

const (
	AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyNewParamsFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyNewParamsFiltersTrafficExclusion string

const (
	AlertingV3PolicyNewParamsFiltersTrafficExclusionSecurityEvents AlertingV3PolicyNewParamsFiltersTrafficExclusion = "security_events"
)

type AlertingV3PolicyUpdateParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertingV3PolicyUpdateParamsAlertType] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[AlertingV3PolicyUpdateParamsFilters] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[interface{}] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r AlertingV3PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyUpdateParamsAlertType string

const (
	AlertingV3PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyUpdateParamsAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedDdosAttackL4Alert                     AlertingV3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedDdosAttackL7Alert                     AlertingV3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyUpdateParamsAlertType = "advanced_http_alert_error"
	AlertingV3PolicyUpdateParamsAlertTypeBgpHijackNotification                         AlertingV3PolicyUpdateParamsAlertType = "bgp_hijack_notification"
	AlertingV3PolicyUpdateParamsAlertTypeBillingUsageAlert                             AlertingV3PolicyUpdateParamsAlertType = "billing_usage_alert"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyUpdateParamsAlertType = "block_notification_block_removed"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyUpdateParamsAlertType = "block_notification_new_block"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyUpdateParamsAlertType = "block_notification_review_rejected"
	AlertingV3PolicyUpdateParamsAlertTypeBrandProtectionAlert                          AlertingV3PolicyUpdateParamsAlertType = "brand_protection_alert"
	AlertingV3PolicyUpdateParamsAlertTypeBrandProtectionDigest                         AlertingV3PolicyUpdateParamsAlertType = "brand_protection_digest"
	AlertingV3PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyUpdateParamsAlertTypeCustomSslCertificateEventType                 AlertingV3PolicyUpdateParamsAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyUpdateParamsAlertTypeDedicatedSslCertificateEventType              AlertingV3PolicyUpdateParamsAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyUpdateParamsAlertTypeDosAttackL4                                   AlertingV3PolicyUpdateParamsAlertType = "dos_attack_l4"
	AlertingV3PolicyUpdateParamsAlertTypeDosAttackL7                                   AlertingV3PolicyUpdateParamsAlertType = "dos_attack_l7"
	AlertingV3PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyUpdateParamsAlertType = "expiring_service_token_alert"
	AlertingV3PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyUpdateParamsAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyUpdateParamsAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyUpdateParamsAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyUpdateParamsAlertTypeFbmDosdAttack                                 AlertingV3PolicyUpdateParamsAlertType = "fbm_dosd_attack"
	AlertingV3PolicyUpdateParamsAlertTypeFbmVolumetricAttack                           AlertingV3PolicyUpdateParamsAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyUpdateParamsAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyUpdateParamsAlertType = "health_check_status_notification"
	AlertingV3PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyUpdateParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyUpdateParamsAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyUpdateParamsAlertType = "http_alert_edge_error"
	AlertingV3PolicyUpdateParamsAlertTypeHTTPAlertOriginError                          AlertingV3PolicyUpdateParamsAlertType = "http_alert_origin_error"
	AlertingV3PolicyUpdateParamsAlertTypeIncidentAlert                                 AlertingV3PolicyUpdateParamsAlertType = "incident_alert"
	AlertingV3PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyUpdateParamsAlertType = "load_balancing_health_alert"
	AlertingV3PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyUpdateParamsAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyUpdateParamsAlertTypeLogoMatchAlert                                AlertingV3PolicyUpdateParamsAlertType = "logo_match_alert"
	AlertingV3PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyUpdateParamsAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyUpdateParamsAlertTypeMaintenanceEventNotification                  AlertingV3PolicyUpdateParamsAlertType = "maintenance_event_notification"
	AlertingV3PolicyUpdateParamsAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyUpdateParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyUpdateParamsAlertTypePagesEventAlert                               AlertingV3PolicyUpdateParamsAlertType = "pages_event_alert"
	AlertingV3PolicyUpdateParamsAlertTypeRadarNotification                             AlertingV3PolicyUpdateParamsAlertType = "radar_notification"
	AlertingV3PolicyUpdateParamsAlertTypeRealOriginMonitoring                          AlertingV3PolicyUpdateParamsAlertType = "real_origin_monitoring"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyUpdateParamsAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyUpdateParamsAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyUpdateParamsAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyUpdateParamsAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyUpdateParamsAlertTypeSentinelAlert                                 AlertingV3PolicyUpdateParamsAlertType = "sentinel_alert"
	AlertingV3PolicyUpdateParamsAlertTypeStreamLiveNotifications                       AlertingV3PolicyUpdateParamsAlertType = "stream_live_notifications"
	AlertingV3PolicyUpdateParamsAlertTypeTunnelHealthEvent                             AlertingV3PolicyUpdateParamsAlertType = "tunnel_health_event"
	AlertingV3PolicyUpdateParamsAlertTypeTunnelUpdateEvent                             AlertingV3PolicyUpdateParamsAlertType = "tunnel_update_event"
	AlertingV3PolicyUpdateParamsAlertTypeUniversalSslEventType                         AlertingV3PolicyUpdateParamsAlertType = "universal_ssl_event_type"
	AlertingV3PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyUpdateParamsAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyUpdateParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyUpdateParamsFilters struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue param.Field[[]AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
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
	IncidentImpact param.Field[[]AlertingV3PolicyUpdateParamsFiltersIncidentImpact] `json:"incident_impact"`
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
	TrafficExclusions param.Field[[]AlertingV3PolicyUpdateParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r AlertingV3PolicyUpdateParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyUpdateParamsFiltersIncidentImpact string

const (
	AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyUpdateParamsFiltersTrafficExclusion string

const (
	AlertingV3PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents AlertingV3PolicyUpdateParamsFiltersTrafficExclusion = "security_events"
)

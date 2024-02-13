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

// Update a Notification policy.
func (r *AlertingV3PolicyService) Update(ctx context.Context, accountID string, policyID string, body AlertingV3PolicyUpdateParams, opts ...option.RequestOption) (res *AlertingV3PolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Notification policy.
func (r *AlertingV3PolicyService) Delete(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AlertingV3PolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single policy.
func (r *AlertingV3PolicyService) Get(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AlertingV3PolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Notification policy.
func (r *AlertingV3PolicyService) NotificationPoliciesNewANotificationPolicy(ctx context.Context, accountID string, body AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams, opts ...option.RequestOption) (res *AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all Notification policies.
func (r *AlertingV3PolicyService) NotificationPoliciesListNotificationPolicies(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingV3PolicyUpdateResponse struct {
	// UUID
	ID   string                             `json:"id"`
	JSON alertingV3PolicyUpdateResponseJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyUpdateResponse]
type alertingV3PolicyUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AlertingV3PolicyDeleteResponseUnknown],
// [AlertingV3PolicyDeleteResponseArray] or [shared.UnionString].
type AlertingV3PolicyDeleteResponse interface {
	ImplementsAlertingV3PolicyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3PolicyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3PolicyDeleteResponseArray []interface{}

func (r AlertingV3PolicyDeleteResponseArray) ImplementsAlertingV3PolicyDeleteResponse() {}

type AlertingV3PolicyGetResponse struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertingV3PolicyGetResponseAlertType `json:"alert_type"`
	Created   time.Time                            `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertingV3PolicyGetResponseFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]AlertingV3PolicyGetResponseMechanisms `json:"mechanisms"`
	Modified   time.Time                                          `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                          `json:"name"`
	JSON alertingV3PolicyGetResponseJSON `json:"-"`
}

// alertingV3PolicyGetResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponse]
type alertingV3PolicyGetResponseJSON struct {
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

func (r *AlertingV3PolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyGetResponseAlertType string

const (
	AlertingV3PolicyGetResponseAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyGetResponseAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyGetResponseAlertTypeAdvancedDDOSAttackL4Alert                     AlertingV3PolicyGetResponseAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyGetResponseAlertTypeAdvancedDDOSAttackL7Alert                     AlertingV3PolicyGetResponseAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyGetResponseAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyGetResponseAlertType = "advanced_http_alert_error"
	AlertingV3PolicyGetResponseAlertTypeBGPHijackNotification                         AlertingV3PolicyGetResponseAlertType = "bgp_hijack_notification"
	AlertingV3PolicyGetResponseAlertTypeBillingUsageAlert                             AlertingV3PolicyGetResponseAlertType = "billing_usage_alert"
	AlertingV3PolicyGetResponseAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyGetResponseAlertType = "block_notification_block_removed"
	AlertingV3PolicyGetResponseAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyGetResponseAlertType = "block_notification_new_block"
	AlertingV3PolicyGetResponseAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyGetResponseAlertType = "block_notification_review_rejected"
	AlertingV3PolicyGetResponseAlertTypeBrandProtectionAlert                          AlertingV3PolicyGetResponseAlertType = "brand_protection_alert"
	AlertingV3PolicyGetResponseAlertTypeBrandProtectionDigest                         AlertingV3PolicyGetResponseAlertType = "brand_protection_digest"
	AlertingV3PolicyGetResponseAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyGetResponseAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyGetResponseAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyGetResponseAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyGetResponseAlertTypeCustomSSLCertificateEventType                 AlertingV3PolicyGetResponseAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyGetResponseAlertTypeDedicatedSSLCertificateEventType              AlertingV3PolicyGetResponseAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyGetResponseAlertTypeDosAttackL4                                   AlertingV3PolicyGetResponseAlertType = "dos_attack_l4"
	AlertingV3PolicyGetResponseAlertTypeDosAttackL7                                   AlertingV3PolicyGetResponseAlertType = "dos_attack_l7"
	AlertingV3PolicyGetResponseAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyGetResponseAlertType = "expiring_service_token_alert"
	AlertingV3PolicyGetResponseAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyGetResponseAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyGetResponseAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyGetResponseAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyGetResponseAlertTypeFbmDosdAttack                                 AlertingV3PolicyGetResponseAlertType = "fbm_dosd_attack"
	AlertingV3PolicyGetResponseAlertTypeFbmVolumetricAttack                           AlertingV3PolicyGetResponseAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyGetResponseAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyGetResponseAlertType = "health_check_status_notification"
	AlertingV3PolicyGetResponseAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyGetResponseAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyGetResponseAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyGetResponseAlertType = "http_alert_edge_error"
	AlertingV3PolicyGetResponseAlertTypeHTTPAlertOriginError                          AlertingV3PolicyGetResponseAlertType = "http_alert_origin_error"
	AlertingV3PolicyGetResponseAlertTypeIncidentAlert                                 AlertingV3PolicyGetResponseAlertType = "incident_alert"
	AlertingV3PolicyGetResponseAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyGetResponseAlertType = "load_balancing_health_alert"
	AlertingV3PolicyGetResponseAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyGetResponseAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyGetResponseAlertTypeLogoMatchAlert                                AlertingV3PolicyGetResponseAlertType = "logo_match_alert"
	AlertingV3PolicyGetResponseAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyGetResponseAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyGetResponseAlertTypeMaintenanceEventNotification                  AlertingV3PolicyGetResponseAlertType = "maintenance_event_notification"
	AlertingV3PolicyGetResponseAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyGetResponseAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyGetResponseAlertTypePagesEventAlert                               AlertingV3PolicyGetResponseAlertType = "pages_event_alert"
	AlertingV3PolicyGetResponseAlertTypeRadarNotification                             AlertingV3PolicyGetResponseAlertType = "radar_notification"
	AlertingV3PolicyGetResponseAlertTypeRealOriginMonitoring                          AlertingV3PolicyGetResponseAlertType = "real_origin_monitoring"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyGetResponseAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyGetResponseAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyGetResponseAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyGetResponseAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyGetResponseAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyGetResponseAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyGetResponseAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyGetResponseAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyGetResponseAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyGetResponseAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyGetResponseAlertTypeSentinelAlert                                 AlertingV3PolicyGetResponseAlertType = "sentinel_alert"
	AlertingV3PolicyGetResponseAlertTypeStreamLiveNotifications                       AlertingV3PolicyGetResponseAlertType = "stream_live_notifications"
	AlertingV3PolicyGetResponseAlertTypeTunnelHealthEvent                             AlertingV3PolicyGetResponseAlertType = "tunnel_health_event"
	AlertingV3PolicyGetResponseAlertTypeTunnelUpdateEvent                             AlertingV3PolicyGetResponseAlertType = "tunnel_update_event"
	AlertingV3PolicyGetResponseAlertTypeUniversalSSLEventType                         AlertingV3PolicyGetResponseAlertType = "universal_ssl_event_type"
	AlertingV3PolicyGetResponseAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyGetResponseAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyGetResponseAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyGetResponseAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyGetResponseFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedAsns []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue []AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []AlertingV3PolicyGetResponseFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []AlertingV3PolicyGetResponseFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                               `json:"zones"`
	JSON  alertingV3PolicyGetResponseFiltersJSON `json:"-"`
}

// alertingV3PolicyGetResponseFiltersJSON contains the JSON metadata for the struct
// [AlertingV3PolicyGetResponseFilters]
type alertingV3PolicyGetResponseFiltersJSON struct {
	Actions                      apijson.Field
	AffectedAsns                 apijson.Field
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

func (r *AlertingV3PolicyGetResponseFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyGetResponseFiltersIncidentImpact string

const (
	AlertingV3PolicyGetResponseFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyGetResponseFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyGetResponseFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyGetResponseFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyGetResponseFiltersTrafficExclusion string

const (
	AlertingV3PolicyGetResponseFiltersTrafficExclusionSecurityEvents AlertingV3PolicyGetResponseFiltersTrafficExclusion = "security_events"
)

type AlertingV3PolicyGetResponseMechanisms struct {
	// UUID
	ID   AlertingV3PolicyGetResponseMechanismsID   `json:"id"`
	JSON alertingV3PolicyGetResponseMechanismsJSON `json:"-"`
}

// alertingV3PolicyGetResponseMechanismsJSON contains the JSON metadata for the
// struct [AlertingV3PolicyGetResponseMechanisms]
type alertingV3PolicyGetResponseMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AlertingV3PolicyGetResponseMechanismsID interface {
	ImplementsAlertingV3PolicyGetResponseMechanismsID()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3PolicyGetResponseMechanismsID)(nil)).Elem(),
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

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponse struct {
	// UUID
	ID   string                                                                 `json:"id"`
	JSON alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseJSON contains
// the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponse]
type alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponse struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType `json:"alert_type"`
	Created   time.Time                                                                     `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanisms `json:"mechanisms"`
	Modified   time.Time                                                                                   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                                                                   `json:"name"`
	JSON alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponse]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseJSON struct {
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

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType string

const (
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeAdvancedDDOSAttackL4Alert                     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeAdvancedDDOSAttackL7Alert                     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "advanced_http_alert_error"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBGPHijackNotification                         AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "bgp_hijack_notification"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBillingUsageAlert                             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "billing_usage_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "block_notification_block_removed"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "block_notification_new_block"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "block_notification_review_rejected"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBrandProtectionAlert                          AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "brand_protection_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeBrandProtectionDigest                         AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "brand_protection_digest"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeCustomSSLCertificateEventType                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeDedicatedSSLCertificateEventType              AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeDosAttackL4                                   AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "dos_attack_l4"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeDosAttackL7                                   AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "dos_attack_l7"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "expiring_service_token_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeFbmDosdAttack                                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "fbm_dosd_attack"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeFbmVolumetricAttack                           AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "health_check_status_notification"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "http_alert_edge_error"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeHTTPAlertOriginError                          AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "http_alert_origin_error"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeIncidentAlert                                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "incident_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "load_balancing_health_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeLogoMatchAlert                                AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "logo_match_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeMaintenanceEventNotification                  AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "maintenance_event_notification"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypePagesEventAlert                               AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "pages_event_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeRadarNotification                             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "radar_notification"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeRealOriginMonitoring                          AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "real_origin_monitoring"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeSentinelAlert                                 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "sentinel_alert"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeStreamLiveNotifications                       AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "stream_live_notifications"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeTunnelHealthEvent                             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "tunnel_health_event"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeTunnelUpdateEvent                             AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "tunnel_update_event"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeUniversalSSLEventType                         AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "universal_ssl_event_type"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedAsns []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                                                                        `json:"zones"`
	JSON  alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFilters]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersJSON struct {
	Actions                      apijson.Field
	AffectedAsns                 apijson.Field
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

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact string

const (
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersTrafficExclusion string

const (
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersTrafficExclusionSecurityEvents AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseFiltersTrafficExclusion = "security_events"
)

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanisms struct {
	// UUID
	ID   AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsID   `json:"id"`
	JSON alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanisms]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsID interface {
	ImplementsAlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsID()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseMechanismsID)(nil)).Elem(),
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
	Mechanisms param.Field[map[string][]AlertingV3PolicyUpdateParamsMechanisms] `json:"mechanisms"`
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
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedDDOSAttackL4Alert                     AlertingV3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedDDOSAttackL7Alert                     AlertingV3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyUpdateParamsAlertType = "advanced_http_alert_error"
	AlertingV3PolicyUpdateParamsAlertTypeBGPHijackNotification                         AlertingV3PolicyUpdateParamsAlertType = "bgp_hijack_notification"
	AlertingV3PolicyUpdateParamsAlertTypeBillingUsageAlert                             AlertingV3PolicyUpdateParamsAlertType = "billing_usage_alert"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyUpdateParamsAlertType = "block_notification_block_removed"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyUpdateParamsAlertType = "block_notification_new_block"
	AlertingV3PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyUpdateParamsAlertType = "block_notification_review_rejected"
	AlertingV3PolicyUpdateParamsAlertTypeBrandProtectionAlert                          AlertingV3PolicyUpdateParamsAlertType = "brand_protection_alert"
	AlertingV3PolicyUpdateParamsAlertTypeBrandProtectionDigest                         AlertingV3PolicyUpdateParamsAlertType = "brand_protection_digest"
	AlertingV3PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType                 AlertingV3PolicyUpdateParamsAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType              AlertingV3PolicyUpdateParamsAlertType = "dedicated_ssl_certificate_event_type"
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
	AlertingV3PolicyUpdateParamsAlertTypeUniversalSSLEventType                         AlertingV3PolicyUpdateParamsAlertType = "universal_ssl_event_type"
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
	AffectedAsns param.Field[[]string] `json:"affected_asns"`
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

type AlertingV3PolicyUpdateParamsMechanisms struct {
	// UUID
	ID param.Field[AlertingV3PolicyUpdateParamsMechanismsID] `json:"id"`
}

func (r AlertingV3PolicyUpdateParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AlertingV3PolicyUpdateParamsMechanismsID interface {
	ImplementsAlertingV3PolicyUpdateParamsMechanismsID()
}

type AlertingV3PolicyUpdateResponseEnvelope struct {
	Errors   []AlertingV3PolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3PolicyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3PolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3PolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// alertingV3PolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AlertingV3PolicyUpdateResponseEnvelope]
type alertingV3PolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    alertingV3PolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AlertingV3PolicyUpdateResponseEnvelopeErrors]
type alertingV3PolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    alertingV3PolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AlertingV3PolicyUpdateResponseEnvelopeMessages]
type alertingV3PolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyUpdateResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyUpdateResponseEnvelopeSuccessTrue AlertingV3PolicyUpdateResponseEnvelopeSuccess = true
)

type AlertingV3PolicyDeleteResponseEnvelope struct {
	Errors   []AlertingV3PolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3PolicyDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3PolicyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3PolicyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3PolicyDeleteResponseEnvelopeJSON       `json:"-"`
}

// alertingV3PolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AlertingV3PolicyDeleteResponseEnvelope]
type alertingV3PolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    alertingV3PolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AlertingV3PolicyDeleteResponseEnvelopeErrors]
type alertingV3PolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    alertingV3PolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AlertingV3PolicyDeleteResponseEnvelopeMessages]
type alertingV3PolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyDeleteResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyDeleteResponseEnvelopeSuccessTrue AlertingV3PolicyDeleteResponseEnvelopeSuccess = true
)

type AlertingV3PolicyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       alertingV3PolicyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3PolicyDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AlertingV3PolicyDeleteResponseEnvelopeResultInfo]
type alertingV3PolicyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseEnvelope struct {
	Errors   []AlertingV3PolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3PolicyGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3PolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3PolicyGetResponseEnvelopeJSON    `json:"-"`
}

// alertingV3PolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AlertingV3PolicyGetResponseEnvelope]
type alertingV3PolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    alertingV3PolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AlertingV3PolicyGetResponseEnvelopeErrors]
type alertingV3PolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    alertingV3PolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AlertingV3PolicyGetResponseEnvelopeMessages]
type alertingV3PolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyGetResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyGetResponseEnvelopeSuccessTrue AlertingV3PolicyGetResponseEnvelopeSuccess = true
)

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanisms] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFilters] `json:"filters"`
}

func (r AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType string

const (
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeAccessCustomCertificateExpirationType         AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "access_custom_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeAdvancedDDOSAttackL4Alert                     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeAdvancedDDOSAttackL7Alert                     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "advanced_http_alert_error"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBGPHijackNotification                         AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "bgp_hijack_notification"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBillingUsageAlert                             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "billing_usage_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "block_notification_block_removed"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "block_notification_new_block"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "block_notification_review_rejected"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBrandProtectionAlert                          AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "brand_protection_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeBrandProtectionDigest                         AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "brand_protection_digest"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeCustomSSLCertificateEventType                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeDedicatedSSLCertificateEventType              AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "dedicated_ssl_certificate_event_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeDosAttackL4                                   AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "dos_attack_l4"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeDosAttackL7                                   AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "dos_attack_l7"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeExpiringServiceTokenAlert                     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "expiring_service_token_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeFailingLogpushJobDisabledAlert                AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "failing_logpush_job_disabled_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeFbmAutoAdvertisement                          AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "fbm_auto_advertisement"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeFbmDosdAttack                                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "fbm_dosd_attack"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeFbmVolumetricAttack                           AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "fbm_volumetric_attack"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeHealthCheckStatusNotification                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "health_check_status_notification"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeHostnameAopCustomCertificateExpirationType    AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeHTTPAlertEdgeError                            AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "http_alert_edge_error"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeHTTPAlertOriginError                          AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "http_alert_origin_error"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeIncidentAlert                                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "incident_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeLoadBalancingHealthAlert                      AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "load_balancing_health_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeLoadBalancingPoolEnablementAlert              AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "load_balancing_pool_enablement_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeLogoMatchAlert                                AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "logo_match_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeMagicTunnelHealthCheckEvent                   AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "magic_tunnel_health_check_event"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeMaintenanceEventNotification                  AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "maintenance_event_notification"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeMtlsCertificateStoreCertificateExpirationType AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypePagesEventAlert                               AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "pages_event_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeRadarNotification                             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "radar_notification"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeRealOriginMonitoring                          AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "real_origin_monitoring"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewHosts                    AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_hosts"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewMaliciousURL             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeScriptmonitorAlertNewResources                AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "scriptmonitor_alert_new_resources"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeSecondaryDNSAllPrimariesFailing               AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "secondary_dns_all_primaries_failing"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeSecondaryDNSPrimariesFailing                  AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "secondary_dns_primaries_failing"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "secondary_dns_zone_successfully_updated"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeSecondaryDNSZoneValidationWarning             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "secondary_dns_zone_validation_warning"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeSentinelAlert                                 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "sentinel_alert"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeStreamLiveNotifications                       AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "stream_live_notifications"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeTunnelHealthEvent                             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "tunnel_health_event"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeTunnelUpdateEvent                             AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "tunnel_update_event"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeUniversalSSLEventType                         AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "universal_ssl_event_type"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanisms struct {
	// UUID
	ID param.Field[AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanismsID] `json:"id"`
}

func (r AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanismsID interface {
	ImplementsAlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsMechanismsID()
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFilters struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedAsns param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Used for configuring magic_tunnel_health_check_event
	AlertTriggerPreferencesValue param.Field[[]AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
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
	IncidentImpact param.Field[[]AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact] `json:"incident_impact"`
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
	TrafficExclusions param.Field[[]AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue string

const (
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue99_0 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue = "99.0"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue98_0 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue = "98.0"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue97_0 AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact string

const (
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpactIncidentImpactNone     AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpactIncidentImpactMinor    AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpactIncidentImpactMajor    AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpactIncidentImpactCritical AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersTrafficExclusion string

const (
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersTrafficExclusionSecurityEvents AlertingV3PolicyNotificationPoliciesNewANotificationPolicyParamsFiltersTrafficExclusion = "security_events"
)

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelope struct {
	Errors   []AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeJSON    `json:"-"`
}

// alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelope]
type alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrors]
type alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessages]
type alertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeSuccessTrue AlertingV3PolicyNotificationPoliciesNewANotificationPolicyResponseEnvelopeSuccess = true
)

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelope struct {
	Errors   []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessages `json:"messages,required"`
	Result   []AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeJSON       `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelope]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrors]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessages]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeSuccessTrue AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeSuccess = true
)

type AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                    `json:"total_count"`
	JSON       alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfo]
type alertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNotificationPoliciesListNotificationPoliciesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

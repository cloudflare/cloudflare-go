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

// Creates a new Notification policy.
func (r *AlertingV3PolicyService) New(ctx context.Context, params AlertingV3PolicyNewParams, opts ...option.RequestOption) (res *AlertingV3PolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Notification policy.
func (r *AlertingV3PolicyService) Update(ctx context.Context, policyID string, params AlertingV3PolicyUpdateParams, opts ...option.RequestOption) (res *AlertingV3PolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all Notification policies.
func (r *AlertingV3PolicyService) List(ctx context.Context, query AlertingV3PolicyListParams, opts ...option.RequestOption) (res *[]AaaPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Notification policy.
func (r *AlertingV3PolicyService) Delete(ctx context.Context, policyID string, body AlertingV3PolicyDeleteParams, opts ...option.RequestOption) (res *AlertingV3PolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single policy.
func (r *AlertingV3PolicyService) Get(ctx context.Context, policyID string, query AlertingV3PolicyGetParams, opts ...option.RequestOption) (res *AaaPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3PolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AaaPolicies struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AaaPoliciesAlertType `json:"alert_type"`
	Created   time.Time            `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AaaPoliciesFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]AaaPoliciesMechanisms `json:"mechanisms"`
	Modified   time.Time                          `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string          `json:"name"`
	JSON aaaPoliciesJSON `json:"-"`
}

// aaaPoliciesJSON contains the JSON metadata for the struct [AaaPolicies]
type aaaPoliciesJSON struct {
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

func (r *AaaPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AaaPoliciesAlertType string

const (
	AaaPoliciesAlertTypeAccessCustomCertificateExpirationType         AaaPoliciesAlertType = "access_custom_certificate_expiration_type"
	AaaPoliciesAlertTypeAdvancedDDOSAttackL4Alert                     AaaPoliciesAlertType = "advanced_ddos_attack_l4_alert"
	AaaPoliciesAlertTypeAdvancedDDOSAttackL7Alert                     AaaPoliciesAlertType = "advanced_ddos_attack_l7_alert"
	AaaPoliciesAlertTypeAdvancedHTTPAlertError                        AaaPoliciesAlertType = "advanced_http_alert_error"
	AaaPoliciesAlertTypeBGPHijackNotification                         AaaPoliciesAlertType = "bgp_hijack_notification"
	AaaPoliciesAlertTypeBillingUsageAlert                             AaaPoliciesAlertType = "billing_usage_alert"
	AaaPoliciesAlertTypeBlockNotificationBlockRemoved                 AaaPoliciesAlertType = "block_notification_block_removed"
	AaaPoliciesAlertTypeBlockNotificationNewBlock                     AaaPoliciesAlertType = "block_notification_new_block"
	AaaPoliciesAlertTypeBlockNotificationReviewRejected               AaaPoliciesAlertType = "block_notification_review_rejected"
	AaaPoliciesAlertTypeBrandProtectionAlert                          AaaPoliciesAlertType = "brand_protection_alert"
	AaaPoliciesAlertTypeBrandProtectionDigest                         AaaPoliciesAlertType = "brand_protection_digest"
	AaaPoliciesAlertTypeClickhouseAlertFwAnomaly                      AaaPoliciesAlertType = "clickhouse_alert_fw_anomaly"
	AaaPoliciesAlertTypeClickhouseAlertFwEntAnomaly                   AaaPoliciesAlertType = "clickhouse_alert_fw_ent_anomaly"
	AaaPoliciesAlertTypeCustomSSLCertificateEventType                 AaaPoliciesAlertType = "custom_ssl_certificate_event_type"
	AaaPoliciesAlertTypeDedicatedSSLCertificateEventType              AaaPoliciesAlertType = "dedicated_ssl_certificate_event_type"
	AaaPoliciesAlertTypeDosAttackL4                                   AaaPoliciesAlertType = "dos_attack_l4"
	AaaPoliciesAlertTypeDosAttackL7                                   AaaPoliciesAlertType = "dos_attack_l7"
	AaaPoliciesAlertTypeExpiringServiceTokenAlert                     AaaPoliciesAlertType = "expiring_service_token_alert"
	AaaPoliciesAlertTypeFailingLogpushJobDisabledAlert                AaaPoliciesAlertType = "failing_logpush_job_disabled_alert"
	AaaPoliciesAlertTypeFbmAutoAdvertisement                          AaaPoliciesAlertType = "fbm_auto_advertisement"
	AaaPoliciesAlertTypeFbmDosdAttack                                 AaaPoliciesAlertType = "fbm_dosd_attack"
	AaaPoliciesAlertTypeFbmVolumetricAttack                           AaaPoliciesAlertType = "fbm_volumetric_attack"
	AaaPoliciesAlertTypeHealthCheckStatusNotification                 AaaPoliciesAlertType = "health_check_status_notification"
	AaaPoliciesAlertTypeHostnameAopCustomCertificateExpirationType    AaaPoliciesAlertType = "hostname_aop_custom_certificate_expiration_type"
	AaaPoliciesAlertTypeHTTPAlertEdgeError                            AaaPoliciesAlertType = "http_alert_edge_error"
	AaaPoliciesAlertTypeHTTPAlertOriginError                          AaaPoliciesAlertType = "http_alert_origin_error"
	AaaPoliciesAlertTypeIncidentAlert                                 AaaPoliciesAlertType = "incident_alert"
	AaaPoliciesAlertTypeLoadBalancingHealthAlert                      AaaPoliciesAlertType = "load_balancing_health_alert"
	AaaPoliciesAlertTypeLoadBalancingPoolEnablementAlert              AaaPoliciesAlertType = "load_balancing_pool_enablement_alert"
	AaaPoliciesAlertTypeLogoMatchAlert                                AaaPoliciesAlertType = "logo_match_alert"
	AaaPoliciesAlertTypeMagicTunnelHealthCheckEvent                   AaaPoliciesAlertType = "magic_tunnel_health_check_event"
	AaaPoliciesAlertTypeMaintenanceEventNotification                  AaaPoliciesAlertType = "maintenance_event_notification"
	AaaPoliciesAlertTypeMTLSCertificateStoreCertificateExpirationType AaaPoliciesAlertType = "mtls_certificate_store_certificate_expiration_type"
	AaaPoliciesAlertTypePagesEventAlert                               AaaPoliciesAlertType = "pages_event_alert"
	AaaPoliciesAlertTypeRadarNotification                             AaaPoliciesAlertType = "radar_notification"
	AaaPoliciesAlertTypeRealOriginMonitoring                          AaaPoliciesAlertType = "real_origin_monitoring"
	AaaPoliciesAlertTypeScriptmonitorAlertNewCodeChangeDetections     AaaPoliciesAlertType = "scriptmonitor_alert_new_code_change_detections"
	AaaPoliciesAlertTypeScriptmonitorAlertNewHosts                    AaaPoliciesAlertType = "scriptmonitor_alert_new_hosts"
	AaaPoliciesAlertTypeScriptmonitorAlertNewMaliciousHosts           AaaPoliciesAlertType = "scriptmonitor_alert_new_malicious_hosts"
	AaaPoliciesAlertTypeScriptmonitorAlertNewMaliciousScripts         AaaPoliciesAlertType = "scriptmonitor_alert_new_malicious_scripts"
	AaaPoliciesAlertTypeScriptmonitorAlertNewMaliciousURL             AaaPoliciesAlertType = "scriptmonitor_alert_new_malicious_url"
	AaaPoliciesAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AaaPoliciesAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AaaPoliciesAlertTypeScriptmonitorAlertNewResources                AaaPoliciesAlertType = "scriptmonitor_alert_new_resources"
	AaaPoliciesAlertTypeSecondaryDNSAllPrimariesFailing               AaaPoliciesAlertType = "secondary_dns_all_primaries_failing"
	AaaPoliciesAlertTypeSecondaryDNSPrimariesFailing                  AaaPoliciesAlertType = "secondary_dns_primaries_failing"
	AaaPoliciesAlertTypeSecondaryDNSZoneSuccessfullyUpdated           AaaPoliciesAlertType = "secondary_dns_zone_successfully_updated"
	AaaPoliciesAlertTypeSecondaryDNSZoneValidationWarning             AaaPoliciesAlertType = "secondary_dns_zone_validation_warning"
	AaaPoliciesAlertTypeSentinelAlert                                 AaaPoliciesAlertType = "sentinel_alert"
	AaaPoliciesAlertTypeStreamLiveNotifications                       AaaPoliciesAlertType = "stream_live_notifications"
	AaaPoliciesAlertTypeTunnelHealthEvent                             AaaPoliciesAlertType = "tunnel_health_event"
	AaaPoliciesAlertTypeTunnelUpdateEvent                             AaaPoliciesAlertType = "tunnel_update_event"
	AaaPoliciesAlertTypeUniversalSSLEventType                         AaaPoliciesAlertType = "universal_ssl_event_type"
	AaaPoliciesAlertTypeWebAnalyticsMetricsUpdate                     AaaPoliciesAlertType = "web_analytics_metrics_update"
	AaaPoliciesAlertTypeZoneAopCustomCertificateExpirationType        AaaPoliciesAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AaaPoliciesFilters struct {
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
	AlertTriggerPreferencesValue []AaaPoliciesFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []AaaPoliciesFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []AaaPoliciesFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string               `json:"zones"`
	JSON  aaaPoliciesFiltersJSON `json:"-"`
}

// aaaPoliciesFiltersJSON contains the JSON metadata for the struct
// [AaaPoliciesFilters]
type aaaPoliciesFiltersJSON struct {
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

func (r *AaaPoliciesFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AaaPoliciesFiltersAlertTriggerPreferencesValue string

const (
	AaaPoliciesFiltersAlertTriggerPreferencesValue99_0 AaaPoliciesFiltersAlertTriggerPreferencesValue = "99.0"
	AaaPoliciesFiltersAlertTriggerPreferencesValue98_0 AaaPoliciesFiltersAlertTriggerPreferencesValue = "98.0"
	AaaPoliciesFiltersAlertTriggerPreferencesValue97_0 AaaPoliciesFiltersAlertTriggerPreferencesValue = "97.0"
)

type AaaPoliciesFiltersIncidentImpact string

const (
	AaaPoliciesFiltersIncidentImpactIncidentImpactNone     AaaPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AaaPoliciesFiltersIncidentImpactIncidentImpactMinor    AaaPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AaaPoliciesFiltersIncidentImpactIncidentImpactMajor    AaaPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AaaPoliciesFiltersIncidentImpactIncidentImpactCritical AaaPoliciesFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type AaaPoliciesFiltersTrafficExclusion string

const (
	AaaPoliciesFiltersTrafficExclusionSecurityEvents AaaPoliciesFiltersTrafficExclusion = "security_events"
)

type AaaPoliciesMechanisms struct {
	// UUID
	ID   AaaPoliciesMechanismsID   `json:"id"`
	JSON aaaPoliciesMechanismsJSON `json:"-"`
}

// aaaPoliciesMechanismsJSON contains the JSON metadata for the struct
// [AaaPoliciesMechanisms]
type aaaPoliciesMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaPoliciesMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AaaPoliciesMechanismsID interface {
	ImplementsAaaPoliciesMechanismsID()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AaaPoliciesMechanismsID)(nil)).Elem(),
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

type AlertingV3PolicyNewResponse struct {
	// UUID
	ID   string                          `json:"id"`
	JSON alertingV3PolicyNewResponseJSON `json:"-"`
}

// alertingV3PolicyNewResponseJSON contains the JSON metadata for the struct
// [AlertingV3PolicyNewResponse]
type alertingV3PolicyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type AlertingV3PolicyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertingV3PolicyNewParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]AlertingV3PolicyNewParamsMechanisms] `json:"mechanisms,required"`
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
	AlertingV3PolicyNewParamsAlertTypeAdvancedDDOSAttackL4Alert                     AlertingV3PolicyNewParamsAlertType = "advanced_ddos_attack_l4_alert"
	AlertingV3PolicyNewParamsAlertTypeAdvancedDDOSAttackL7Alert                     AlertingV3PolicyNewParamsAlertType = "advanced_ddos_attack_l7_alert"
	AlertingV3PolicyNewParamsAlertTypeAdvancedHTTPAlertError                        AlertingV3PolicyNewParamsAlertType = "advanced_http_alert_error"
	AlertingV3PolicyNewParamsAlertTypeBGPHijackNotification                         AlertingV3PolicyNewParamsAlertType = "bgp_hijack_notification"
	AlertingV3PolicyNewParamsAlertTypeBillingUsageAlert                             AlertingV3PolicyNewParamsAlertType = "billing_usage_alert"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationBlockRemoved                 AlertingV3PolicyNewParamsAlertType = "block_notification_block_removed"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationNewBlock                     AlertingV3PolicyNewParamsAlertType = "block_notification_new_block"
	AlertingV3PolicyNewParamsAlertTypeBlockNotificationReviewRejected               AlertingV3PolicyNewParamsAlertType = "block_notification_review_rejected"
	AlertingV3PolicyNewParamsAlertTypeBrandProtectionAlert                          AlertingV3PolicyNewParamsAlertType = "brand_protection_alert"
	AlertingV3PolicyNewParamsAlertTypeBrandProtectionDigest                         AlertingV3PolicyNewParamsAlertType = "brand_protection_digest"
	AlertingV3PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly                      AlertingV3PolicyNewParamsAlertType = "clickhouse_alert_fw_anomaly"
	AlertingV3PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly                   AlertingV3PolicyNewParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertingV3PolicyNewParamsAlertTypeCustomSSLCertificateEventType                 AlertingV3PolicyNewParamsAlertType = "custom_ssl_certificate_event_type"
	AlertingV3PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType              AlertingV3PolicyNewParamsAlertType = "dedicated_ssl_certificate_event_type"
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
	AlertingV3PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType AlertingV3PolicyNewParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
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
	AlertingV3PolicyNewParamsAlertTypeUniversalSSLEventType                         AlertingV3PolicyNewParamsAlertType = "universal_ssl_event_type"
	AlertingV3PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate                     AlertingV3PolicyNewParamsAlertType = "web_analytics_metrics_update"
	AlertingV3PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType        AlertingV3PolicyNewParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

type AlertingV3PolicyNewParamsMechanisms struct {
	// UUID
	ID param.Field[AlertingV3PolicyNewParamsMechanismsID] `json:"id"`
}

func (r AlertingV3PolicyNewParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AlertingV3PolicyNewParamsMechanismsID interface {
	ImplementsAlertingV3PolicyNewParamsMechanismsID()
}

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

type AlertingV3PolicyNewResponseEnvelope struct {
	Errors   []AlertingV3PolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3PolicyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3PolicyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3PolicyNewResponseEnvelopeJSON    `json:"-"`
}

// alertingV3PolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AlertingV3PolicyNewResponseEnvelope]
type alertingV3PolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    alertingV3PolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AlertingV3PolicyNewResponseEnvelopeErrors]
type alertingV3PolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    alertingV3PolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AlertingV3PolicyNewResponseEnvelopeMessages]
type alertingV3PolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyNewResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyNewResponseEnvelopeSuccessTrue AlertingV3PolicyNewResponseEnvelopeSuccess = true
)

type AlertingV3PolicyUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
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
	AlertingV3PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType AlertingV3PolicyUpdateParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
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

type AlertingV3PolicyListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3PolicyListResponseEnvelope struct {
	Errors   []AlertingV3PolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AaaPolicies                                  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3PolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3PolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3PolicyListResponseEnvelopeJSON       `json:"-"`
}

// alertingV3PolicyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AlertingV3PolicyListResponseEnvelope]
type alertingV3PolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    alertingV3PolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3PolicyListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AlertingV3PolicyListResponseEnvelopeErrors]
type alertingV3PolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    alertingV3PolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3PolicyListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AlertingV3PolicyListResponseEnvelopeMessages]
type alertingV3PolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3PolicyListResponseEnvelopeSuccess bool

const (
	AlertingV3PolicyListResponseEnvelopeSuccessTrue AlertingV3PolicyListResponseEnvelopeSuccess = true
)

type AlertingV3PolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       alertingV3PolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3PolicyListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AlertingV3PolicyListResponseEnvelopeResultInfo]
type alertingV3PolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3PolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3PolicyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

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

type AlertingV3PolicyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3PolicyGetResponseEnvelope struct {
	Errors   []AlertingV3PolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3PolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AaaPolicies                                   `json:"result,required"`
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

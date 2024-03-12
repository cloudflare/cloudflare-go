// File generated from our OpenAPI spec by Stainless.

package alerting

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// V3PolicyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewV3PolicyService] method instead.
type V3PolicyService struct {
	Options []option.RequestOption
}

// NewV3PolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3PolicyService(opts ...option.RequestOption) (r *V3PolicyService) {
	r = &V3PolicyService{}
	r.Options = opts
	return
}

// Creates a new Notification policy.
func (r *V3PolicyService) New(ctx context.Context, params V3PolicyNewParams, opts ...option.RequestOption) (res *V3PolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3PolicyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Notification policy.
func (r *V3PolicyService) Update(ctx context.Context, policyID string, params V3PolicyUpdateParams, opts ...option.RequestOption) (res *V3PolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3PolicyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all Notification policies.
func (r *V3PolicyService) List(ctx context.Context, query V3PolicyListParams, opts ...option.RequestOption) (res *[]V3PolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3PolicyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Notification policy.
func (r *V3PolicyService) Delete(ctx context.Context, policyID string, body V3PolicyDeleteParams, opts ...option.RequestOption) (res *V3PolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3PolicyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single policy.
func (r *V3PolicyService) Get(ctx context.Context, policyID string, query V3PolicyGetParams, opts ...option.RequestOption) (res *V3PolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3PolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V3PolicyNewResponse struct {
	// UUID
	ID   string                  `json:"id"`
	JSON v3PolicyNewResponseJSON `json:"-"`
}

// v3PolicyNewResponseJSON contains the JSON metadata for the struct
// [V3PolicyNewResponse]
type v3PolicyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type V3PolicyUpdateResponse struct {
	// UUID
	ID   string                     `json:"id"`
	JSON v3PolicyUpdateResponseJSON `json:"-"`
}

// v3PolicyUpdateResponseJSON contains the JSON metadata for the struct
// [V3PolicyUpdateResponse]
type v3PolicyUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V3PolicyListResponse struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType V3PolicyListResponseAlertType `json:"alert_type"`
	Created   time.Time                     `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters V3PolicyListResponseFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]V3PolicyListResponseMechanisms `json:"mechanisms"`
	Modified   time.Time                                   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                   `json:"name"`
	JSON v3PolicyListResponseJSON `json:"-"`
}

// v3PolicyListResponseJSON contains the JSON metadata for the struct
// [V3PolicyListResponse]
type v3PolicyListResponseJSON struct {
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

func (r *V3PolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type V3PolicyListResponseAlertType string

const (
	V3PolicyListResponseAlertTypeAccessCustomCertificateExpirationType         V3PolicyListResponseAlertType = "access_custom_certificate_expiration_type"
	V3PolicyListResponseAlertTypeAdvancedDDOSAttackL4Alert                     V3PolicyListResponseAlertType = "advanced_ddos_attack_l4_alert"
	V3PolicyListResponseAlertTypeAdvancedDDOSAttackL7Alert                     V3PolicyListResponseAlertType = "advanced_ddos_attack_l7_alert"
	V3PolicyListResponseAlertTypeAdvancedHTTPAlertError                        V3PolicyListResponseAlertType = "advanced_http_alert_error"
	V3PolicyListResponseAlertTypeBGPHijackNotification                         V3PolicyListResponseAlertType = "bgp_hijack_notification"
	V3PolicyListResponseAlertTypeBillingUsageAlert                             V3PolicyListResponseAlertType = "billing_usage_alert"
	V3PolicyListResponseAlertTypeBlockNotificationBlockRemoved                 V3PolicyListResponseAlertType = "block_notification_block_removed"
	V3PolicyListResponseAlertTypeBlockNotificationNewBlock                     V3PolicyListResponseAlertType = "block_notification_new_block"
	V3PolicyListResponseAlertTypeBlockNotificationReviewRejected               V3PolicyListResponseAlertType = "block_notification_review_rejected"
	V3PolicyListResponseAlertTypeBrandProtectionAlert                          V3PolicyListResponseAlertType = "brand_protection_alert"
	V3PolicyListResponseAlertTypeBrandProtectionDigest                         V3PolicyListResponseAlertType = "brand_protection_digest"
	V3PolicyListResponseAlertTypeClickhouseAlertFwAnomaly                      V3PolicyListResponseAlertType = "clickhouse_alert_fw_anomaly"
	V3PolicyListResponseAlertTypeClickhouseAlertFwEntAnomaly                   V3PolicyListResponseAlertType = "clickhouse_alert_fw_ent_anomaly"
	V3PolicyListResponseAlertTypeCustomSSLCertificateEventType                 V3PolicyListResponseAlertType = "custom_ssl_certificate_event_type"
	V3PolicyListResponseAlertTypeDedicatedSSLCertificateEventType              V3PolicyListResponseAlertType = "dedicated_ssl_certificate_event_type"
	V3PolicyListResponseAlertTypeDosAttackL4                                   V3PolicyListResponseAlertType = "dos_attack_l4"
	V3PolicyListResponseAlertTypeDosAttackL7                                   V3PolicyListResponseAlertType = "dos_attack_l7"
	V3PolicyListResponseAlertTypeExpiringServiceTokenAlert                     V3PolicyListResponseAlertType = "expiring_service_token_alert"
	V3PolicyListResponseAlertTypeFailingLogpushJobDisabledAlert                V3PolicyListResponseAlertType = "failing_logpush_job_disabled_alert"
	V3PolicyListResponseAlertTypeFbmAutoAdvertisement                          V3PolicyListResponseAlertType = "fbm_auto_advertisement"
	V3PolicyListResponseAlertTypeFbmDosdAttack                                 V3PolicyListResponseAlertType = "fbm_dosd_attack"
	V3PolicyListResponseAlertTypeFbmVolumetricAttack                           V3PolicyListResponseAlertType = "fbm_volumetric_attack"
	V3PolicyListResponseAlertTypeHealthCheckStatusNotification                 V3PolicyListResponseAlertType = "health_check_status_notification"
	V3PolicyListResponseAlertTypeHostnameAopCustomCertificateExpirationType    V3PolicyListResponseAlertType = "hostname_aop_custom_certificate_expiration_type"
	V3PolicyListResponseAlertTypeHTTPAlertEdgeError                            V3PolicyListResponseAlertType = "http_alert_edge_error"
	V3PolicyListResponseAlertTypeHTTPAlertOriginError                          V3PolicyListResponseAlertType = "http_alert_origin_error"
	V3PolicyListResponseAlertTypeIncidentAlert                                 V3PolicyListResponseAlertType = "incident_alert"
	V3PolicyListResponseAlertTypeLoadBalancingHealthAlert                      V3PolicyListResponseAlertType = "load_balancing_health_alert"
	V3PolicyListResponseAlertTypeLoadBalancingPoolEnablementAlert              V3PolicyListResponseAlertType = "load_balancing_pool_enablement_alert"
	V3PolicyListResponseAlertTypeLogoMatchAlert                                V3PolicyListResponseAlertType = "logo_match_alert"
	V3PolicyListResponseAlertTypeMagicTunnelHealthCheckEvent                   V3PolicyListResponseAlertType = "magic_tunnel_health_check_event"
	V3PolicyListResponseAlertTypeMaintenanceEventNotification                  V3PolicyListResponseAlertType = "maintenance_event_notification"
	V3PolicyListResponseAlertTypeMTLSCertificateStoreCertificateExpirationType V3PolicyListResponseAlertType = "mtls_certificate_store_certificate_expiration_type"
	V3PolicyListResponseAlertTypePagesEventAlert                               V3PolicyListResponseAlertType = "pages_event_alert"
	V3PolicyListResponseAlertTypeRadarNotification                             V3PolicyListResponseAlertType = "radar_notification"
	V3PolicyListResponseAlertTypeRealOriginMonitoring                          V3PolicyListResponseAlertType = "real_origin_monitoring"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewCodeChangeDetections     V3PolicyListResponseAlertType = "scriptmonitor_alert_new_code_change_detections"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewHosts                    V3PolicyListResponseAlertType = "scriptmonitor_alert_new_hosts"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewMaliciousHosts           V3PolicyListResponseAlertType = "scriptmonitor_alert_new_malicious_hosts"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewMaliciousScripts         V3PolicyListResponseAlertType = "scriptmonitor_alert_new_malicious_scripts"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewMaliciousURL             V3PolicyListResponseAlertType = "scriptmonitor_alert_new_malicious_url"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     V3PolicyListResponseAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	V3PolicyListResponseAlertTypeScriptmonitorAlertNewResources                V3PolicyListResponseAlertType = "scriptmonitor_alert_new_resources"
	V3PolicyListResponseAlertTypeSecondaryDNSAllPrimariesFailing               V3PolicyListResponseAlertType = "secondary_dns_all_primaries_failing"
	V3PolicyListResponseAlertTypeSecondaryDNSPrimariesFailing                  V3PolicyListResponseAlertType = "secondary_dns_primaries_failing"
	V3PolicyListResponseAlertTypeSecondaryDNSZoneSuccessfullyUpdated           V3PolicyListResponseAlertType = "secondary_dns_zone_successfully_updated"
	V3PolicyListResponseAlertTypeSecondaryDNSZoneValidationWarning             V3PolicyListResponseAlertType = "secondary_dns_zone_validation_warning"
	V3PolicyListResponseAlertTypeSentinelAlert                                 V3PolicyListResponseAlertType = "sentinel_alert"
	V3PolicyListResponseAlertTypeStreamLiveNotifications                       V3PolicyListResponseAlertType = "stream_live_notifications"
	V3PolicyListResponseAlertTypeTrafficAnomaliesAlert                         V3PolicyListResponseAlertType = "traffic_anomalies_alert"
	V3PolicyListResponseAlertTypeTunnelHealthEvent                             V3PolicyListResponseAlertType = "tunnel_health_event"
	V3PolicyListResponseAlertTypeTunnelUpdateEvent                             V3PolicyListResponseAlertType = "tunnel_update_event"
	V3PolicyListResponseAlertTypeUniversalSSLEventType                         V3PolicyListResponseAlertType = "universal_ssl_event_type"
	V3PolicyListResponseAlertTypeWebAnalyticsMetricsUpdate                     V3PolicyListResponseAlertType = "web_analytics_metrics_update"
	V3PolicyListResponseAlertTypeZoneAopCustomCertificateExpirationType        V3PolicyListResponseAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type V3PolicyListResponseFilters struct {
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
	AlertTriggerPreferencesValue []V3PolicyListResponseFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []V3PolicyListResponseFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []V3PolicyListResponseFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                        `json:"zones"`
	JSON  v3PolicyListResponseFiltersJSON `json:"-"`
}

// v3PolicyListResponseFiltersJSON contains the JSON metadata for the struct
// [V3PolicyListResponseFilters]
type v3PolicyListResponseFiltersJSON struct {
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

func (r *V3PolicyListResponseFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseFiltersJSON) RawJSON() string {
	return r.raw
}

type V3PolicyListResponseFiltersAlertTriggerPreferencesValue string

const (
	V3PolicyListResponseFiltersAlertTriggerPreferencesValue99_0 V3PolicyListResponseFiltersAlertTriggerPreferencesValue = "99.0"
	V3PolicyListResponseFiltersAlertTriggerPreferencesValue98_0 V3PolicyListResponseFiltersAlertTriggerPreferencesValue = "98.0"
	V3PolicyListResponseFiltersAlertTriggerPreferencesValue97_0 V3PolicyListResponseFiltersAlertTriggerPreferencesValue = "97.0"
)

type V3PolicyListResponseFiltersIncidentImpact string

const (
	V3PolicyListResponseFiltersIncidentImpactIncidentImpactNone     V3PolicyListResponseFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	V3PolicyListResponseFiltersIncidentImpactIncidentImpactMinor    V3PolicyListResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	V3PolicyListResponseFiltersIncidentImpactIncidentImpactMajor    V3PolicyListResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	V3PolicyListResponseFiltersIncidentImpactIncidentImpactCritical V3PolicyListResponseFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type V3PolicyListResponseFiltersTrafficExclusion string

const (
	V3PolicyListResponseFiltersTrafficExclusionSecurityEvents V3PolicyListResponseFiltersTrafficExclusion = "security_events"
)

type V3PolicyListResponseMechanisms struct {
	// UUID
	ID   V3PolicyListResponseMechanismsID   `json:"id"`
	JSON v3PolicyListResponseMechanismsJSON `json:"-"`
}

// v3PolicyListResponseMechanismsJSON contains the JSON metadata for the struct
// [V3PolicyListResponseMechanisms]
type v3PolicyListResponseMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyListResponseMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseMechanismsJSON) RawJSON() string {
	return r.raw
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type V3PolicyListResponseMechanismsID interface {
	ImplementsAlertingV3PolicyListResponseMechanismsID()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3PolicyListResponseMechanismsID)(nil)).Elem(),
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

// Union satisfied by [alerting.V3PolicyDeleteResponseUnknown],
// [alerting.V3PolicyDeleteResponseArray] or [shared.UnionString].
type V3PolicyDeleteResponse interface {
	ImplementsAlertingV3PolicyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3PolicyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3PolicyDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V3PolicyDeleteResponseArray []interface{}

func (r V3PolicyDeleteResponseArray) ImplementsAlertingV3PolicyDeleteResponse() {}

type V3PolicyGetResponse struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType V3PolicyGetResponseAlertType `json:"alert_type"`
	Created   time.Time                    `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters V3PolicyGetResponseFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms map[string][]V3PolicyGetResponseMechanisms `json:"mechanisms"`
	Modified   time.Time                                  `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                  `json:"name"`
	JSON v3PolicyGetResponseJSON `json:"-"`
}

// v3PolicyGetResponseJSON contains the JSON metadata for the struct
// [V3PolicyGetResponse]
type v3PolicyGetResponseJSON struct {
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

func (r *V3PolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type V3PolicyGetResponseAlertType string

const (
	V3PolicyGetResponseAlertTypeAccessCustomCertificateExpirationType         V3PolicyGetResponseAlertType = "access_custom_certificate_expiration_type"
	V3PolicyGetResponseAlertTypeAdvancedDDOSAttackL4Alert                     V3PolicyGetResponseAlertType = "advanced_ddos_attack_l4_alert"
	V3PolicyGetResponseAlertTypeAdvancedDDOSAttackL7Alert                     V3PolicyGetResponseAlertType = "advanced_ddos_attack_l7_alert"
	V3PolicyGetResponseAlertTypeAdvancedHTTPAlertError                        V3PolicyGetResponseAlertType = "advanced_http_alert_error"
	V3PolicyGetResponseAlertTypeBGPHijackNotification                         V3PolicyGetResponseAlertType = "bgp_hijack_notification"
	V3PolicyGetResponseAlertTypeBillingUsageAlert                             V3PolicyGetResponseAlertType = "billing_usage_alert"
	V3PolicyGetResponseAlertTypeBlockNotificationBlockRemoved                 V3PolicyGetResponseAlertType = "block_notification_block_removed"
	V3PolicyGetResponseAlertTypeBlockNotificationNewBlock                     V3PolicyGetResponseAlertType = "block_notification_new_block"
	V3PolicyGetResponseAlertTypeBlockNotificationReviewRejected               V3PolicyGetResponseAlertType = "block_notification_review_rejected"
	V3PolicyGetResponseAlertTypeBrandProtectionAlert                          V3PolicyGetResponseAlertType = "brand_protection_alert"
	V3PolicyGetResponseAlertTypeBrandProtectionDigest                         V3PolicyGetResponseAlertType = "brand_protection_digest"
	V3PolicyGetResponseAlertTypeClickhouseAlertFwAnomaly                      V3PolicyGetResponseAlertType = "clickhouse_alert_fw_anomaly"
	V3PolicyGetResponseAlertTypeClickhouseAlertFwEntAnomaly                   V3PolicyGetResponseAlertType = "clickhouse_alert_fw_ent_anomaly"
	V3PolicyGetResponseAlertTypeCustomSSLCertificateEventType                 V3PolicyGetResponseAlertType = "custom_ssl_certificate_event_type"
	V3PolicyGetResponseAlertTypeDedicatedSSLCertificateEventType              V3PolicyGetResponseAlertType = "dedicated_ssl_certificate_event_type"
	V3PolicyGetResponseAlertTypeDosAttackL4                                   V3PolicyGetResponseAlertType = "dos_attack_l4"
	V3PolicyGetResponseAlertTypeDosAttackL7                                   V3PolicyGetResponseAlertType = "dos_attack_l7"
	V3PolicyGetResponseAlertTypeExpiringServiceTokenAlert                     V3PolicyGetResponseAlertType = "expiring_service_token_alert"
	V3PolicyGetResponseAlertTypeFailingLogpushJobDisabledAlert                V3PolicyGetResponseAlertType = "failing_logpush_job_disabled_alert"
	V3PolicyGetResponseAlertTypeFbmAutoAdvertisement                          V3PolicyGetResponseAlertType = "fbm_auto_advertisement"
	V3PolicyGetResponseAlertTypeFbmDosdAttack                                 V3PolicyGetResponseAlertType = "fbm_dosd_attack"
	V3PolicyGetResponseAlertTypeFbmVolumetricAttack                           V3PolicyGetResponseAlertType = "fbm_volumetric_attack"
	V3PolicyGetResponseAlertTypeHealthCheckStatusNotification                 V3PolicyGetResponseAlertType = "health_check_status_notification"
	V3PolicyGetResponseAlertTypeHostnameAopCustomCertificateExpirationType    V3PolicyGetResponseAlertType = "hostname_aop_custom_certificate_expiration_type"
	V3PolicyGetResponseAlertTypeHTTPAlertEdgeError                            V3PolicyGetResponseAlertType = "http_alert_edge_error"
	V3PolicyGetResponseAlertTypeHTTPAlertOriginError                          V3PolicyGetResponseAlertType = "http_alert_origin_error"
	V3PolicyGetResponseAlertTypeIncidentAlert                                 V3PolicyGetResponseAlertType = "incident_alert"
	V3PolicyGetResponseAlertTypeLoadBalancingHealthAlert                      V3PolicyGetResponseAlertType = "load_balancing_health_alert"
	V3PolicyGetResponseAlertTypeLoadBalancingPoolEnablementAlert              V3PolicyGetResponseAlertType = "load_balancing_pool_enablement_alert"
	V3PolicyGetResponseAlertTypeLogoMatchAlert                                V3PolicyGetResponseAlertType = "logo_match_alert"
	V3PolicyGetResponseAlertTypeMagicTunnelHealthCheckEvent                   V3PolicyGetResponseAlertType = "magic_tunnel_health_check_event"
	V3PolicyGetResponseAlertTypeMaintenanceEventNotification                  V3PolicyGetResponseAlertType = "maintenance_event_notification"
	V3PolicyGetResponseAlertTypeMTLSCertificateStoreCertificateExpirationType V3PolicyGetResponseAlertType = "mtls_certificate_store_certificate_expiration_type"
	V3PolicyGetResponseAlertTypePagesEventAlert                               V3PolicyGetResponseAlertType = "pages_event_alert"
	V3PolicyGetResponseAlertTypeRadarNotification                             V3PolicyGetResponseAlertType = "radar_notification"
	V3PolicyGetResponseAlertTypeRealOriginMonitoring                          V3PolicyGetResponseAlertType = "real_origin_monitoring"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewCodeChangeDetections     V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_code_change_detections"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewHosts                    V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_hosts"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousHosts           V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_hosts"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousScripts         V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_scripts"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaliciousURL             V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_malicious_url"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	V3PolicyGetResponseAlertTypeScriptmonitorAlertNewResources                V3PolicyGetResponseAlertType = "scriptmonitor_alert_new_resources"
	V3PolicyGetResponseAlertTypeSecondaryDNSAllPrimariesFailing               V3PolicyGetResponseAlertType = "secondary_dns_all_primaries_failing"
	V3PolicyGetResponseAlertTypeSecondaryDNSPrimariesFailing                  V3PolicyGetResponseAlertType = "secondary_dns_primaries_failing"
	V3PolicyGetResponseAlertTypeSecondaryDNSZoneSuccessfullyUpdated           V3PolicyGetResponseAlertType = "secondary_dns_zone_successfully_updated"
	V3PolicyGetResponseAlertTypeSecondaryDNSZoneValidationWarning             V3PolicyGetResponseAlertType = "secondary_dns_zone_validation_warning"
	V3PolicyGetResponseAlertTypeSentinelAlert                                 V3PolicyGetResponseAlertType = "sentinel_alert"
	V3PolicyGetResponseAlertTypeStreamLiveNotifications                       V3PolicyGetResponseAlertType = "stream_live_notifications"
	V3PolicyGetResponseAlertTypeTrafficAnomaliesAlert                         V3PolicyGetResponseAlertType = "traffic_anomalies_alert"
	V3PolicyGetResponseAlertTypeTunnelHealthEvent                             V3PolicyGetResponseAlertType = "tunnel_health_event"
	V3PolicyGetResponseAlertTypeTunnelUpdateEvent                             V3PolicyGetResponseAlertType = "tunnel_update_event"
	V3PolicyGetResponseAlertTypeUniversalSSLEventType                         V3PolicyGetResponseAlertType = "universal_ssl_event_type"
	V3PolicyGetResponseAlertTypeWebAnalyticsMetricsUpdate                     V3PolicyGetResponseAlertType = "web_analytics_metrics_update"
	V3PolicyGetResponseAlertTypeZoneAopCustomCertificateExpirationType        V3PolicyGetResponseAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type V3PolicyGetResponseFilters struct {
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
	AlertTriggerPreferencesValue []V3PolicyGetResponseFiltersAlertTriggerPreferencesValue `json:"alert_trigger_preferences_value"`
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
	IncidentImpact []V3PolicyGetResponseFiltersIncidentImpact `json:"incident_impact"`
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
	TrafficExclusions []V3PolicyGetResponseFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string                       `json:"zones"`
	JSON  v3PolicyGetResponseFiltersJSON `json:"-"`
}

// v3PolicyGetResponseFiltersJSON contains the JSON metadata for the struct
// [V3PolicyGetResponseFilters]
type v3PolicyGetResponseFiltersJSON struct {
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

func (r *V3PolicyGetResponseFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseFiltersJSON) RawJSON() string {
	return r.raw
}

type V3PolicyGetResponseFiltersAlertTriggerPreferencesValue string

const (
	V3PolicyGetResponseFiltersAlertTriggerPreferencesValue99_0 V3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "99.0"
	V3PolicyGetResponseFiltersAlertTriggerPreferencesValue98_0 V3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "98.0"
	V3PolicyGetResponseFiltersAlertTriggerPreferencesValue97_0 V3PolicyGetResponseFiltersAlertTriggerPreferencesValue = "97.0"
)

type V3PolicyGetResponseFiltersIncidentImpact string

const (
	V3PolicyGetResponseFiltersIncidentImpactIncidentImpactNone     V3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	V3PolicyGetResponseFiltersIncidentImpactIncidentImpactMinor    V3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	V3PolicyGetResponseFiltersIncidentImpactIncidentImpactMajor    V3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	V3PolicyGetResponseFiltersIncidentImpactIncidentImpactCritical V3PolicyGetResponseFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type V3PolicyGetResponseFiltersTrafficExclusion string

const (
	V3PolicyGetResponseFiltersTrafficExclusionSecurityEvents V3PolicyGetResponseFiltersTrafficExclusion = "security_events"
)

type V3PolicyGetResponseMechanisms struct {
	// UUID
	ID   V3PolicyGetResponseMechanismsID   `json:"id"`
	JSON v3PolicyGetResponseMechanismsJSON `json:"-"`
}

// v3PolicyGetResponseMechanismsJSON contains the JSON metadata for the struct
// [V3PolicyGetResponseMechanisms]
type v3PolicyGetResponseMechanismsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyGetResponseMechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseMechanismsJSON) RawJSON() string {
	return r.raw
}

// UUID
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type V3PolicyGetResponseMechanismsID interface {
	ImplementsAlertingV3PolicyGetResponseMechanismsID()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3PolicyGetResponseMechanismsID)(nil)).Elem(),
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

type V3PolicyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[V3PolicyNewParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]V3PolicyNewParamsMechanisms] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[V3PolicyNewParamsFilters] `json:"filters"`
}

func (r V3PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type V3PolicyNewParamsAlertType string

const (
	V3PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType         V3PolicyNewParamsAlertType = "access_custom_certificate_expiration_type"
	V3PolicyNewParamsAlertTypeAdvancedDDOSAttackL4Alert                     V3PolicyNewParamsAlertType = "advanced_ddos_attack_l4_alert"
	V3PolicyNewParamsAlertTypeAdvancedDDOSAttackL7Alert                     V3PolicyNewParamsAlertType = "advanced_ddos_attack_l7_alert"
	V3PolicyNewParamsAlertTypeAdvancedHTTPAlertError                        V3PolicyNewParamsAlertType = "advanced_http_alert_error"
	V3PolicyNewParamsAlertTypeBGPHijackNotification                         V3PolicyNewParamsAlertType = "bgp_hijack_notification"
	V3PolicyNewParamsAlertTypeBillingUsageAlert                             V3PolicyNewParamsAlertType = "billing_usage_alert"
	V3PolicyNewParamsAlertTypeBlockNotificationBlockRemoved                 V3PolicyNewParamsAlertType = "block_notification_block_removed"
	V3PolicyNewParamsAlertTypeBlockNotificationNewBlock                     V3PolicyNewParamsAlertType = "block_notification_new_block"
	V3PolicyNewParamsAlertTypeBlockNotificationReviewRejected               V3PolicyNewParamsAlertType = "block_notification_review_rejected"
	V3PolicyNewParamsAlertTypeBrandProtectionAlert                          V3PolicyNewParamsAlertType = "brand_protection_alert"
	V3PolicyNewParamsAlertTypeBrandProtectionDigest                         V3PolicyNewParamsAlertType = "brand_protection_digest"
	V3PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly                      V3PolicyNewParamsAlertType = "clickhouse_alert_fw_anomaly"
	V3PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly                   V3PolicyNewParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	V3PolicyNewParamsAlertTypeCustomSSLCertificateEventType                 V3PolicyNewParamsAlertType = "custom_ssl_certificate_event_type"
	V3PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType              V3PolicyNewParamsAlertType = "dedicated_ssl_certificate_event_type"
	V3PolicyNewParamsAlertTypeDosAttackL4                                   V3PolicyNewParamsAlertType = "dos_attack_l4"
	V3PolicyNewParamsAlertTypeDosAttackL7                                   V3PolicyNewParamsAlertType = "dos_attack_l7"
	V3PolicyNewParamsAlertTypeExpiringServiceTokenAlert                     V3PolicyNewParamsAlertType = "expiring_service_token_alert"
	V3PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert                V3PolicyNewParamsAlertType = "failing_logpush_job_disabled_alert"
	V3PolicyNewParamsAlertTypeFbmAutoAdvertisement                          V3PolicyNewParamsAlertType = "fbm_auto_advertisement"
	V3PolicyNewParamsAlertTypeFbmDosdAttack                                 V3PolicyNewParamsAlertType = "fbm_dosd_attack"
	V3PolicyNewParamsAlertTypeFbmVolumetricAttack                           V3PolicyNewParamsAlertType = "fbm_volumetric_attack"
	V3PolicyNewParamsAlertTypeHealthCheckStatusNotification                 V3PolicyNewParamsAlertType = "health_check_status_notification"
	V3PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType    V3PolicyNewParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	V3PolicyNewParamsAlertTypeHTTPAlertEdgeError                            V3PolicyNewParamsAlertType = "http_alert_edge_error"
	V3PolicyNewParamsAlertTypeHTTPAlertOriginError                          V3PolicyNewParamsAlertType = "http_alert_origin_error"
	V3PolicyNewParamsAlertTypeIncidentAlert                                 V3PolicyNewParamsAlertType = "incident_alert"
	V3PolicyNewParamsAlertTypeLoadBalancingHealthAlert                      V3PolicyNewParamsAlertType = "load_balancing_health_alert"
	V3PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert              V3PolicyNewParamsAlertType = "load_balancing_pool_enablement_alert"
	V3PolicyNewParamsAlertTypeLogoMatchAlert                                V3PolicyNewParamsAlertType = "logo_match_alert"
	V3PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent                   V3PolicyNewParamsAlertType = "magic_tunnel_health_check_event"
	V3PolicyNewParamsAlertTypeMaintenanceEventNotification                  V3PolicyNewParamsAlertType = "maintenance_event_notification"
	V3PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType V3PolicyNewParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	V3PolicyNewParamsAlertTypePagesEventAlert                               V3PolicyNewParamsAlertType = "pages_event_alert"
	V3PolicyNewParamsAlertTypeRadarNotification                             V3PolicyNewParamsAlertType = "radar_notification"
	V3PolicyNewParamsAlertTypeRealOriginMonitoring                          V3PolicyNewParamsAlertType = "real_origin_monitoring"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts                    V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_hosts"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL             V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	V3PolicyNewParamsAlertTypeScriptmonitorAlertNewResources                V3PolicyNewParamsAlertType = "scriptmonitor_alert_new_resources"
	V3PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing               V3PolicyNewParamsAlertType = "secondary_dns_all_primaries_failing"
	V3PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing                  V3PolicyNewParamsAlertType = "secondary_dns_primaries_failing"
	V3PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           V3PolicyNewParamsAlertType = "secondary_dns_zone_successfully_updated"
	V3PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning             V3PolicyNewParamsAlertType = "secondary_dns_zone_validation_warning"
	V3PolicyNewParamsAlertTypeSentinelAlert                                 V3PolicyNewParamsAlertType = "sentinel_alert"
	V3PolicyNewParamsAlertTypeStreamLiveNotifications                       V3PolicyNewParamsAlertType = "stream_live_notifications"
	V3PolicyNewParamsAlertTypeTrafficAnomaliesAlert                         V3PolicyNewParamsAlertType = "traffic_anomalies_alert"
	V3PolicyNewParamsAlertTypeTunnelHealthEvent                             V3PolicyNewParamsAlertType = "tunnel_health_event"
	V3PolicyNewParamsAlertTypeTunnelUpdateEvent                             V3PolicyNewParamsAlertType = "tunnel_update_event"
	V3PolicyNewParamsAlertTypeUniversalSSLEventType                         V3PolicyNewParamsAlertType = "universal_ssl_event_type"
	V3PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate                     V3PolicyNewParamsAlertType = "web_analytics_metrics_update"
	V3PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType        V3PolicyNewParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

type V3PolicyNewParamsMechanisms struct {
	// UUID
	ID param.Field[V3PolicyNewParamsMechanismsID] `json:"id"`
}

func (r V3PolicyNewParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type V3PolicyNewParamsMechanismsID interface {
	ImplementsAlertingV3PolicyNewParamsMechanismsID()
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type V3PolicyNewParamsFilters struct {
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
	AlertTriggerPreferencesValue param.Field[[]V3PolicyNewParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
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
	IncidentImpact param.Field[[]V3PolicyNewParamsFiltersIncidentImpact] `json:"incident_impact"`
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
	TrafficExclusions param.Field[[]V3PolicyNewParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r V3PolicyNewParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V3PolicyNewParamsFiltersAlertTriggerPreferencesValue string

const (
	V3PolicyNewParamsFiltersAlertTriggerPreferencesValue99_0 V3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "99.0"
	V3PolicyNewParamsFiltersAlertTriggerPreferencesValue98_0 V3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "98.0"
	V3PolicyNewParamsFiltersAlertTriggerPreferencesValue97_0 V3PolicyNewParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

type V3PolicyNewParamsFiltersIncidentImpact string

const (
	V3PolicyNewParamsFiltersIncidentImpactIncidentImpactNone     V3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	V3PolicyNewParamsFiltersIncidentImpactIncidentImpactMinor    V3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	V3PolicyNewParamsFiltersIncidentImpactIncidentImpactMajor    V3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	V3PolicyNewParamsFiltersIncidentImpactIncidentImpactCritical V3PolicyNewParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type V3PolicyNewParamsFiltersTrafficExclusion string

const (
	V3PolicyNewParamsFiltersTrafficExclusionSecurityEvents V3PolicyNewParamsFiltersTrafficExclusion = "security_events"
)

type V3PolicyNewResponseEnvelope struct {
	Errors   []V3PolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3PolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   V3PolicyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3PolicyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3PolicyNewResponseEnvelopeJSON    `json:"-"`
}

// v3PolicyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3PolicyNewResponseEnvelope]
type v3PolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3PolicyNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v3PolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// v3PolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V3PolicyNewResponseEnvelopeErrors]
type v3PolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3PolicyNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v3PolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// v3PolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V3PolicyNewResponseEnvelopeMessages]
type v3PolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3PolicyNewResponseEnvelopeSuccess bool

const (
	V3PolicyNewResponseEnvelopeSuccessTrue V3PolicyNewResponseEnvelopeSuccess = true
)

type V3PolicyUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[V3PolicyUpdateParamsAlertType] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[V3PolicyUpdateParamsFilters] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[map[string][]V3PolicyUpdateParamsMechanisms] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r V3PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type V3PolicyUpdateParamsAlertType string

const (
	V3PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType         V3PolicyUpdateParamsAlertType = "access_custom_certificate_expiration_type"
	V3PolicyUpdateParamsAlertTypeAdvancedDDOSAttackL4Alert                     V3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l4_alert"
	V3PolicyUpdateParamsAlertTypeAdvancedDDOSAttackL7Alert                     V3PolicyUpdateParamsAlertType = "advanced_ddos_attack_l7_alert"
	V3PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError                        V3PolicyUpdateParamsAlertType = "advanced_http_alert_error"
	V3PolicyUpdateParamsAlertTypeBGPHijackNotification                         V3PolicyUpdateParamsAlertType = "bgp_hijack_notification"
	V3PolicyUpdateParamsAlertTypeBillingUsageAlert                             V3PolicyUpdateParamsAlertType = "billing_usage_alert"
	V3PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved                 V3PolicyUpdateParamsAlertType = "block_notification_block_removed"
	V3PolicyUpdateParamsAlertTypeBlockNotificationNewBlock                     V3PolicyUpdateParamsAlertType = "block_notification_new_block"
	V3PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected               V3PolicyUpdateParamsAlertType = "block_notification_review_rejected"
	V3PolicyUpdateParamsAlertTypeBrandProtectionAlert                          V3PolicyUpdateParamsAlertType = "brand_protection_alert"
	V3PolicyUpdateParamsAlertTypeBrandProtectionDigest                         V3PolicyUpdateParamsAlertType = "brand_protection_digest"
	V3PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly                      V3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_anomaly"
	V3PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly                   V3PolicyUpdateParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	V3PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType                 V3PolicyUpdateParamsAlertType = "custom_ssl_certificate_event_type"
	V3PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType              V3PolicyUpdateParamsAlertType = "dedicated_ssl_certificate_event_type"
	V3PolicyUpdateParamsAlertTypeDosAttackL4                                   V3PolicyUpdateParamsAlertType = "dos_attack_l4"
	V3PolicyUpdateParamsAlertTypeDosAttackL7                                   V3PolicyUpdateParamsAlertType = "dos_attack_l7"
	V3PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert                     V3PolicyUpdateParamsAlertType = "expiring_service_token_alert"
	V3PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert                V3PolicyUpdateParamsAlertType = "failing_logpush_job_disabled_alert"
	V3PolicyUpdateParamsAlertTypeFbmAutoAdvertisement                          V3PolicyUpdateParamsAlertType = "fbm_auto_advertisement"
	V3PolicyUpdateParamsAlertTypeFbmDosdAttack                                 V3PolicyUpdateParamsAlertType = "fbm_dosd_attack"
	V3PolicyUpdateParamsAlertTypeFbmVolumetricAttack                           V3PolicyUpdateParamsAlertType = "fbm_volumetric_attack"
	V3PolicyUpdateParamsAlertTypeHealthCheckStatusNotification                 V3PolicyUpdateParamsAlertType = "health_check_status_notification"
	V3PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType    V3PolicyUpdateParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	V3PolicyUpdateParamsAlertTypeHTTPAlertEdgeError                            V3PolicyUpdateParamsAlertType = "http_alert_edge_error"
	V3PolicyUpdateParamsAlertTypeHTTPAlertOriginError                          V3PolicyUpdateParamsAlertType = "http_alert_origin_error"
	V3PolicyUpdateParamsAlertTypeIncidentAlert                                 V3PolicyUpdateParamsAlertType = "incident_alert"
	V3PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert                      V3PolicyUpdateParamsAlertType = "load_balancing_health_alert"
	V3PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert              V3PolicyUpdateParamsAlertType = "load_balancing_pool_enablement_alert"
	V3PolicyUpdateParamsAlertTypeLogoMatchAlert                                V3PolicyUpdateParamsAlertType = "logo_match_alert"
	V3PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent                   V3PolicyUpdateParamsAlertType = "magic_tunnel_health_check_event"
	V3PolicyUpdateParamsAlertTypeMaintenanceEventNotification                  V3PolicyUpdateParamsAlertType = "maintenance_event_notification"
	V3PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType V3PolicyUpdateParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	V3PolicyUpdateParamsAlertTypePagesEventAlert                               V3PolicyUpdateParamsAlertType = "pages_event_alert"
	V3PolicyUpdateParamsAlertTypeRadarNotification                             V3PolicyUpdateParamsAlertType = "radar_notification"
	V3PolicyUpdateParamsAlertTypeRealOriginMonitoring                          V3PolicyUpdateParamsAlertType = "real_origin_monitoring"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts                    V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_hosts"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL             V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	V3PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources                V3PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_resources"
	V3PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing               V3PolicyUpdateParamsAlertType = "secondary_dns_all_primaries_failing"
	V3PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing                  V3PolicyUpdateParamsAlertType = "secondary_dns_primaries_failing"
	V3PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           V3PolicyUpdateParamsAlertType = "secondary_dns_zone_successfully_updated"
	V3PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning             V3PolicyUpdateParamsAlertType = "secondary_dns_zone_validation_warning"
	V3PolicyUpdateParamsAlertTypeSentinelAlert                                 V3PolicyUpdateParamsAlertType = "sentinel_alert"
	V3PolicyUpdateParamsAlertTypeStreamLiveNotifications                       V3PolicyUpdateParamsAlertType = "stream_live_notifications"
	V3PolicyUpdateParamsAlertTypeTrafficAnomaliesAlert                         V3PolicyUpdateParamsAlertType = "traffic_anomalies_alert"
	V3PolicyUpdateParamsAlertTypeTunnelHealthEvent                             V3PolicyUpdateParamsAlertType = "tunnel_health_event"
	V3PolicyUpdateParamsAlertTypeTunnelUpdateEvent                             V3PolicyUpdateParamsAlertType = "tunnel_update_event"
	V3PolicyUpdateParamsAlertTypeUniversalSSLEventType                         V3PolicyUpdateParamsAlertType = "universal_ssl_event_type"
	V3PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate                     V3PolicyUpdateParamsAlertType = "web_analytics_metrics_update"
	V3PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType        V3PolicyUpdateParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type V3PolicyUpdateParamsFilters struct {
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
	AlertTriggerPreferencesValue param.Field[[]V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue] `json:"alert_trigger_preferences_value"`
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
	IncidentImpact param.Field[[]V3PolicyUpdateParamsFiltersIncidentImpact] `json:"incident_impact"`
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
	TrafficExclusions param.Field[[]V3PolicyUpdateParamsFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Used for configuring magic_tunnel_health_check_event
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r V3PolicyUpdateParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue string

const (
	V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue99_0 V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "99.0"
	V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue98_0 V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "98.0"
	V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue97_0 V3PolicyUpdateParamsFiltersAlertTriggerPreferencesValue = "97.0"
)

type V3PolicyUpdateParamsFiltersIncidentImpact string

const (
	V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactNone     V3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMinor    V3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactMajor    V3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	V3PolicyUpdateParamsFiltersIncidentImpactIncidentImpactCritical V3PolicyUpdateParamsFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

type V3PolicyUpdateParamsFiltersTrafficExclusion string

const (
	V3PolicyUpdateParamsFiltersTrafficExclusionSecurityEvents V3PolicyUpdateParamsFiltersTrafficExclusion = "security_events"
)

type V3PolicyUpdateParamsMechanisms struct {
	// UUID
	ID param.Field[V3PolicyUpdateParamsMechanismsID] `json:"id"`
}

func (r V3PolicyUpdateParamsMechanisms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UUID
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type V3PolicyUpdateParamsMechanismsID interface {
	ImplementsAlertingV3PolicyUpdateParamsMechanismsID()
}

type V3PolicyUpdateResponseEnvelope struct {
	Errors   []V3PolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3PolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   V3PolicyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3PolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3PolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// v3PolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3PolicyUpdateResponseEnvelope]
type v3PolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3PolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v3PolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// v3PolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V3PolicyUpdateResponseEnvelopeErrors]
type v3PolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3PolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    v3PolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// v3PolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V3PolicyUpdateResponseEnvelopeMessages]
type v3PolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3PolicyUpdateResponseEnvelopeSuccess bool

const (
	V3PolicyUpdateResponseEnvelopeSuccessTrue V3PolicyUpdateResponseEnvelopeSuccess = true
)

type V3PolicyListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3PolicyListResponseEnvelope struct {
	Errors   []V3PolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3PolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []V3PolicyListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3PolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3PolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3PolicyListResponseEnvelopeJSON       `json:"-"`
}

// v3PolicyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3PolicyListResponseEnvelope]
type v3PolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3PolicyListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    v3PolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// v3PolicyListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V3PolicyListResponseEnvelopeErrors]
type v3PolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3PolicyListResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v3PolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// v3PolicyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V3PolicyListResponseEnvelopeMessages]
type v3PolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3PolicyListResponseEnvelopeSuccess bool

const (
	V3PolicyListResponseEnvelopeSuccessTrue V3PolicyListResponseEnvelopeSuccess = true
)

type V3PolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       v3PolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3PolicyListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [V3PolicyListResponseEnvelopeResultInfo]
type v3PolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3PolicyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3PolicyDeleteResponseEnvelope struct {
	Errors   []V3PolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3PolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V3PolicyDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3PolicyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3PolicyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3PolicyDeleteResponseEnvelopeJSON       `json:"-"`
}

// v3PolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3PolicyDeleteResponseEnvelope]
type v3PolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3PolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v3PolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v3PolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V3PolicyDeleteResponseEnvelopeErrors]
type v3PolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3PolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    v3PolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v3PolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V3PolicyDeleteResponseEnvelopeMessages]
type v3PolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3PolicyDeleteResponseEnvelopeSuccess bool

const (
	V3PolicyDeleteResponseEnvelopeSuccessTrue V3PolicyDeleteResponseEnvelopeSuccess = true
)

type V3PolicyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       v3PolicyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3PolicyDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [V3PolicyDeleteResponseEnvelopeResultInfo]
type v3PolicyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3PolicyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3PolicyGetResponseEnvelope struct {
	Errors   []V3PolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3PolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   V3PolicyGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3PolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3PolicyGetResponseEnvelopeJSON    `json:"-"`
}

// v3PolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3PolicyGetResponseEnvelope]
type v3PolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3PolicyGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v3PolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v3PolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V3PolicyGetResponseEnvelopeErrors]
type v3PolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3PolicyGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v3PolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v3PolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V3PolicyGetResponseEnvelopeMessages]
type v3PolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3PolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3PolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3PolicyGetResponseEnvelopeSuccess bool

const (
	V3PolicyGetResponseEnvelopeSuccessTrue V3PolicyGetResponseEnvelopeSuccess = true
)

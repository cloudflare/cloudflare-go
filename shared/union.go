// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsUserAuditLogListParamsBeforeUnion()      {}
func (UnionTime) ImplementsUserAuditLogListParamsSinceUnion()       {}
func (UnionTime) ImplementsAuditLogsAuditLogListParamsBeforeUnion() {}
func (UnionTime) ImplementsAuditLogsAuditLogListParamsSinceUnion()  {}

type UnionString string

func (UnionString) ImplementsFirewallWAFPackageListResponseFirewallAPIResponseCollectionResultUnion() {
}
func (UnionString) ImplementsLogsReceivedGetParamsEndUnion()                 {}
func (UnionString) ImplementsLogsReceivedGetParamsStartUnion()               {}
func (UnionString) ImplementsWorkersAIRunResponseUnion()                     {}
func (UnionString) ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion() {}
func (UnionString) ImplementsQueuesQueueDeleteResponseUnion()                {}
func (UnionString) ImplementsQueuesConsumerDeleteResponseUnion()             {}
func (UnionString) ImplementsSpectrumOriginPortUnionParam()                  {}
func (UnionString) ImplementsSpectrumOriginPortUnion()                       {}
func (UnionString) ImplementsMagicTransitCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitCfInterconnectListResponseInterconnectsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitCfInterconnectGetResponseInterconnectHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitCfInterconnectUpdateParamsHealthCheckTargetUnion()     {}
func (UnionString) ImplementsMagicTransitGRETunnelNewResponseGRETunnelsHealthCheckTargetUnion() {}
func (UnionString) ImplementsMagicTransitGRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitGRETunnelListResponseGRETunnelsHealthCheckTargetUnion() {}
func (UnionString) ImplementsMagicTransitGRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitGRETunnelGetResponseGRETunnelHealthCheckTargetUnion()      {}
func (UnionString) ImplementsMagicTransitGRETunnelUpdateParamsHealthCheckTargetUnion()              {}
func (UnionString) ImplementsMagicTransitIPSECTunnelNewResponseIPSECTunnelsHealthCheckTargetUnion() {}
func (UnionString) ImplementsMagicTransitIPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitIPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitIPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion() {}
func (UnionString) ImplementsMagicTransitIPSECTunnelNewParamsHealthCheckTargetUnion()              {}
func (UnionString) ImplementsMagicTransitIPSECTunnelUpdateParamsHealthCheckTargetUnion()           {}
func (UnionString) ImplementsRegistrarDomainUpdateResponseUnion()                                  {}
func (UnionString) ImplementsRegistrarDomainListResponseResultUnion()                              {}
func (UnionString) ImplementsRegistrarDomainGetResponseUnion()                                     {}
func (UnionString) ImplementsRulesListItemGetResponseUnion()                                       {}
func (UnionString) ImplementsWARPConnectorWARPConnectorTokenResponseUnion()                        {}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodySaaSApplicationPolicyUnion() {}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion() {}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsZeroTrustTunnelTokenGetResponseUnion()            {}
func (UnionString) ImplementsZeroTrustTunnelManagementNewResponseUnion()       {}
func (UnionString) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union() {}
func (UnionString) ImplementsHostnamesSettingValueUnionParam()                 {}
func (UnionString) ImplementsHostnamesSettingValueUnion()                      {}
func (UnionString) ImplementsAIGatewayLogListParamsFiltersValueUnion()         {}
func (UnionString) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()       {}

type UnionBool bool

func (UnionBool) ImplementsAIGatewayLogListParamsFiltersValueUnion()   {}
func (UnionBool) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion() {}

type UnionInt int64

func (UnionInt) ImplementsLogsReceivedGetParamsEndUnion()   {}
func (UnionInt) ImplementsLogsReceivedGetParamsStartUnion() {}
func (UnionInt) ImplementsSpectrumOriginPortUnionParam()    {}
func (UnionInt) ImplementsSpectrumOriginPortUnion()         {}
func (UnionInt) ImplementsRulesListItemGetResponseUnion()   {}

type UnionFloat float64

func (UnionFloat) ImplementsDNSTTL()                                          {}
func (UnionFloat) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union() {}
func (UnionFloat) ImplementsHostnamesSettingValueUnionParam()                 {}
func (UnionFloat) ImplementsHostnamesSettingValueUnion()                      {}
func (UnionFloat) ImplementsAIGatewayLogListParamsFiltersValueUnion()         {}
func (UnionFloat) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()       {}

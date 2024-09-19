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
func (UnionString) ImplementsRegistrarDomainUpdateResponseUnion()            {}
func (UnionString) ImplementsRegistrarDomainListResponseResultUnion()        {}
func (UnionString) ImplementsRegistrarDomainGetResponseUnion()               {}
func (UnionString) ImplementsRulesListItemGetResponseUnion()                 {}
func (UnionString) ImplementsWARPConnectorWARPConnectorTokenResponseUnion()  {}
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
func (UnionString) ImplementsZeroTrustTunnelTokenGetResponseUnion()                      {}
func (UnionString) ImplementsZeroTrustTunnelManagementNewResponseUnion()                 {}
func (UnionString) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()           {}
func (UnionString) ImplementsHostnamesSettingValueUnionParam()                           {}
func (UnionString) ImplementsHostnamesSettingValueUnion()                                {}
func (UnionString) ImplementsEventNotificationsR2ConfigurationQueueDeleteResponseUnion() {}

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

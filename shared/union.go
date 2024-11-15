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

func (UnionString) ImplementsLogsReceivedGetParamsEndUnion()                                    {}
func (UnionString) ImplementsLogsReceivedGetParamsStartUnion()                                  {}
func (UnionString) ImplementsWorkersAIRunResponseUnion()                                        {}
func (UnionString) ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion()                    {}
func (UnionString) ImplementsSpectrumOriginPortUnionParam()                                     {}
func (UnionString) ImplementsSpectrumOriginPortUnion()                                          {}
func (UnionString) ImplementsMagicTransitHealthCheckTargetUnionParam()                          {}
func (UnionString) ImplementsMagicTransitHealthCheckTargetUnion()                               {}
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
func (UnionString) ImplementsRulesListItemListResponseUnion()                                      {}
func (UnionString) ImplementsRulesListItemGetResponseUnion()                                       {}
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
func (UnionString) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()               {}
func (UnionString) ImplementsHostnamesSettingValueUnionParam()                               {}
func (UnionString) ImplementsHostnamesSettingValueUnion()                                    {}
func (UnionString) ImplementsAIGatewayLogListParamsFiltersValueUnion()                       {}
func (UnionString) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()                     {}
func (UnionString) ImplementsAIGatewayLogEditParamsMetadataUnion()                           {}
func (UnionString) ImplementsAIGatewayDatasetNewResponseFiltersValueUnion()                  {}
func (UnionString) ImplementsAIGatewayDatasetUpdateResponseFiltersValueUnion()               {}
func (UnionString) ImplementsAIGatewayDatasetListResponseFiltersValueUnion()                 {}
func (UnionString) ImplementsAIGatewayDatasetDeleteResponseFiltersValueUnion()               {}
func (UnionString) ImplementsAIGatewayDatasetGetResponseFiltersValueUnion()                  {}
func (UnionString) ImplementsAIGatewayDatasetNewParamsFiltersValueUnion()                    {}
func (UnionString) ImplementsAIGatewayDatasetUpdateParamsFiltersValueUnion()                 {}
func (UnionString) ImplementsAIGatewayEvaluationNewResponseDatasetsFiltersValueUnion()       {}
func (UnionString) ImplementsAIGatewayEvaluationListResponseDatasetsFiltersValueUnion()      {}
func (UnionString) ImplementsAIGatewayEvaluationDeleteResponseDatasetsFiltersValueUnion()    {}
func (UnionString) ImplementsAIGatewayEvaluationGetResponseDatasetsFiltersValueUnion()       {}
func (UnionString) ImplementsWorkflowInstanceGetResponseOutputUnion()                        {}
func (UnionString) ImplementsWorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion() {}
func (UnionString) ImplementsWorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion()      {}

type UnionBool bool

func (UnionBool) ImplementsAIGatewayLogListParamsFiltersValueUnion()                    {}
func (UnionBool) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()                  {}
func (UnionBool) ImplementsAIGatewayLogEditParamsMetadataUnion()                        {}
func (UnionBool) ImplementsAIGatewayDatasetNewResponseFiltersValueUnion()               {}
func (UnionBool) ImplementsAIGatewayDatasetUpdateResponseFiltersValueUnion()            {}
func (UnionBool) ImplementsAIGatewayDatasetListResponseFiltersValueUnion()              {}
func (UnionBool) ImplementsAIGatewayDatasetDeleteResponseFiltersValueUnion()            {}
func (UnionBool) ImplementsAIGatewayDatasetGetResponseFiltersValueUnion()               {}
func (UnionBool) ImplementsAIGatewayDatasetNewParamsFiltersValueUnion()                 {}
func (UnionBool) ImplementsAIGatewayDatasetUpdateParamsFiltersValueUnion()              {}
func (UnionBool) ImplementsAIGatewayEvaluationNewResponseDatasetsFiltersValueUnion()    {}
func (UnionBool) ImplementsAIGatewayEvaluationListResponseDatasetsFiltersValueUnion()   {}
func (UnionBool) ImplementsAIGatewayEvaluationDeleteResponseDatasetsFiltersValueUnion() {}
func (UnionBool) ImplementsAIGatewayEvaluationGetResponseDatasetsFiltersValueUnion()    {}

type UnionInt int64

func (UnionInt) ImplementsLogsReceivedGetParamsEndUnion()   {}
func (UnionInt) ImplementsLogsReceivedGetParamsStartUnion() {}
func (UnionInt) ImplementsSpectrumOriginPortUnionParam()    {}
func (UnionInt) ImplementsSpectrumOriginPortUnion()         {}
func (UnionInt) ImplementsRulesListItemListResponseUnion()  {}
func (UnionInt) ImplementsRulesListItemGetResponseUnion()   {}

type UnionFloat float64

func (UnionFloat) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()               {}
func (UnionFloat) ImplementsHostnamesSettingValueUnionParam()                               {}
func (UnionFloat) ImplementsHostnamesSettingValueUnion()                                    {}
func (UnionFloat) ImplementsAIGatewayLogListParamsFiltersValueUnion()                       {}
func (UnionFloat) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsAIGatewayLogEditParamsMetadataUnion()                           {}
func (UnionFloat) ImplementsAIGatewayDatasetNewResponseFiltersValueUnion()                  {}
func (UnionFloat) ImplementsAIGatewayDatasetUpdateResponseFiltersValueUnion()               {}
func (UnionFloat) ImplementsAIGatewayDatasetListResponseFiltersValueUnion()                 {}
func (UnionFloat) ImplementsAIGatewayDatasetDeleteResponseFiltersValueUnion()               {}
func (UnionFloat) ImplementsAIGatewayDatasetGetResponseFiltersValueUnion()                  {}
func (UnionFloat) ImplementsAIGatewayDatasetNewParamsFiltersValueUnion()                    {}
func (UnionFloat) ImplementsAIGatewayDatasetUpdateParamsFiltersValueUnion()                 {}
func (UnionFloat) ImplementsAIGatewayEvaluationNewResponseDatasetsFiltersValueUnion()       {}
func (UnionFloat) ImplementsAIGatewayEvaluationListResponseDatasetsFiltersValueUnion()      {}
func (UnionFloat) ImplementsAIGatewayEvaluationDeleteResponseDatasetsFiltersValueUnion()    {}
func (UnionFloat) ImplementsAIGatewayEvaluationGetResponseDatasetsFiltersValueUnion()       {}
func (UnionFloat) ImplementsWorkflowInstanceGetResponseOutputUnion()                        {}
func (UnionFloat) ImplementsWorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion() {}
func (UnionFloat) ImplementsWorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion()      {}

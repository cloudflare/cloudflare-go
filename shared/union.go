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
func (UnionString) ImplementsMagicTransitGRETunnelBulkUpdateResponseModifiedGRETunnelsHealthCheckTargetUnion() {
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
func (UnionString) ImplementsMagicTransitIPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsMagicTransitIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion() {}
func (UnionString) ImplementsMagicTransitIPSECTunnelNewParamsHealthCheckTargetUnion()              {}
func (UnionString) ImplementsMagicTransitIPSECTunnelUpdateParamsHealthCheckTargetUnion()           {}
func (UnionString) ImplementsRulesListItemListResponseUnion()                                      {}
func (UnionString) ImplementsRulesListItemGetResponseUnion()                                       {}
func (UnionString) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()                     {}
func (UnionString) ImplementsHostnamesSettingValueUnionParam()                                     {}
func (UnionString) ImplementsHostnamesSettingValueUnion()                                          {}
func (UnionString) ImplementsAIGatewayLogListParamsFiltersValueUnion()                             {}
func (UnionString) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()                           {}
func (UnionString) ImplementsAIGatewayLogEditParamsMetadataUnion()                                 {}
func (UnionString) ImplementsAIGatewayDatasetNewResponseFiltersValueUnion()                        {}
func (UnionString) ImplementsAIGatewayDatasetUpdateResponseFiltersValueUnion()                     {}
func (UnionString) ImplementsAIGatewayDatasetListResponseFiltersValueUnion()                       {}
func (UnionString) ImplementsAIGatewayDatasetDeleteResponseFiltersValueUnion()                     {}
func (UnionString) ImplementsAIGatewayDatasetGetResponseFiltersValueUnion()                        {}
func (UnionString) ImplementsAIGatewayDatasetNewParamsFiltersValueUnion()                          {}
func (UnionString) ImplementsAIGatewayDatasetUpdateParamsFiltersValueUnion()                       {}
func (UnionString) ImplementsAIGatewayEvaluationNewResponseDatasetsFiltersValueUnion()             {}
func (UnionString) ImplementsAIGatewayEvaluationListResponseDatasetsFiltersValueUnion()            {}
func (UnionString) ImplementsAIGatewayEvaluationDeleteResponseDatasetsFiltersValueUnion()          {}
func (UnionString) ImplementsAIGatewayEvaluationGetResponseDatasetsFiltersValueUnion()             {}
func (UnionString) ImplementsWorkflowsInstanceGetResponseOutputUnion()                             {}
func (UnionString) ImplementsWorkflowsInstanceGetResponseStepsObjectConfigRetriesDelayUnion()      {}
func (UnionString) ImplementsWorkflowsInstanceGetResponseStepsObjectConfigTimeoutUnion()           {}
func (UnionString) ImplementsAIAIRunResponseUnion()                                                {}
func (UnionString) ImplementsAIAIRunParamsBodyTextEmbeddingsTextUnion()                            {}

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

func (UnionInt) ImplementsLogsReceivedGetParamsEndUnion()                                  {}
func (UnionInt) ImplementsLogsReceivedGetParamsStartUnion()                                {}
func (UnionInt) ImplementsPageRulesPageRuleActionsCacheTTLByStatusValueUnion()             {}
func (UnionInt) ImplementsPageRulesPageRuleNewParamsActionsCacheTTLByStatusValueUnion()    {}
func (UnionInt) ImplementsPageRulesPageRuleUpdateParamsActionsCacheTTLByStatusValueUnion() {}
func (UnionInt) ImplementsPageRulesPageRuleEditParamsActionsCacheTTLByStatusValueUnion()   {}
func (UnionInt) ImplementsSpectrumOriginPortUnionParam()                                   {}
func (UnionInt) ImplementsSpectrumOriginPortUnion()                                        {}
func (UnionInt) ImplementsRulesListItemListResponseUnion()                                 {}
func (UnionInt) ImplementsRulesListItemGetResponseUnion()                                  {}

type UnionFloat float64

func (UnionFloat) ImplementsDNSTTL()                                                         {}
func (UnionFloat) ImplementsRadarRankingTimeseriesGroupsResponseSerie0Union()                {}
func (UnionFloat) ImplementsHostnamesSettingValueUnionParam()                                {}
func (UnionFloat) ImplementsHostnamesSettingValueUnion()                                     {}
func (UnionFloat) ImplementsAIGatewayLogListParamsFiltersValueUnion()                        {}
func (UnionFloat) ImplementsAIGatewayLogDeleteParamsFiltersValueUnion()                      {}
func (UnionFloat) ImplementsAIGatewayLogEditParamsMetadataUnion()                            {}
func (UnionFloat) ImplementsAIGatewayDatasetNewResponseFiltersValueUnion()                   {}
func (UnionFloat) ImplementsAIGatewayDatasetUpdateResponseFiltersValueUnion()                {}
func (UnionFloat) ImplementsAIGatewayDatasetListResponseFiltersValueUnion()                  {}
func (UnionFloat) ImplementsAIGatewayDatasetDeleteResponseFiltersValueUnion()                {}
func (UnionFloat) ImplementsAIGatewayDatasetGetResponseFiltersValueUnion()                   {}
func (UnionFloat) ImplementsAIGatewayDatasetNewParamsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsAIGatewayDatasetUpdateParamsFiltersValueUnion()                  {}
func (UnionFloat) ImplementsAIGatewayEvaluationNewResponseDatasetsFiltersValueUnion()        {}
func (UnionFloat) ImplementsAIGatewayEvaluationListResponseDatasetsFiltersValueUnion()       {}
func (UnionFloat) ImplementsAIGatewayEvaluationDeleteResponseDatasetsFiltersValueUnion()     {}
func (UnionFloat) ImplementsAIGatewayEvaluationGetResponseDatasetsFiltersValueUnion()        {}
func (UnionFloat) ImplementsWorkflowsInstanceGetResponseOutputUnion()                        {}
func (UnionFloat) ImplementsWorkflowsInstanceGetResponseStepsObjectConfigRetriesDelayUnion() {}
func (UnionFloat) ImplementsWorkflowsInstanceGetResponseStepsObjectConfigTimeoutUnion()      {}

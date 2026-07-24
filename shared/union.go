// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsAuditLogListParamsBeforeUnion() {}
func (UnionTime) ImplementsAuditLogListParamsSinceUnion()  {}

type UnionString string

func (UnionString) ImplementsAnalyticsQuerySummaryParamsFiltersValueUnion()                        {}
func (UnionString) ImplementsAnalyticsQueryTimeseriesParamsFiltersValueUnion()                     {}
func (UnionString) ImplementsAnalyticsQueryTopNParamsFiltersValueUnion()                           {}
func (UnionString) ImplementsAnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion() {}
func (UnionString) ImplementsAnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion()     {}
func (UnionString) ImplementsAnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion()  {}
func (UnionString) ImplementsEmailSendingSendParamsFromUnion()                                     {}
func (UnionString) ImplementsEmailSendingSendParamsBccUnion()                                      {}
func (UnionString) ImplementsEmailSendingSendParamsBccArrayItemUnion()                             {}
func (UnionString) ImplementsEmailSendingSendParamsCcUnion()                                       {}
func (UnionString) ImplementsEmailSendingSendParamsCcArrayItemUnion()                              {}
func (UnionString) ImplementsEmailSendingSendParamsReplyToUnion()                                  {}
func (UnionString) ImplementsEmailSendingSendParamsToUnion()                                       {}
func (UnionString) ImplementsEmailSendingSendParamsToArrayItemUnion()                              {}
func (UnionString) ImplementsReceivedGetParamsEndUnion()                                           {}
func (UnionString) ImplementsReceivedGetParamsStartUnion()                                         {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsEventsSourceUnion()           {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()       {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseInvocationsSourceUnion()            {}
func (UnionString) ImplementsObservabilityTelemetryValuesResponseValueUnion()                       {}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()    {}
func (UnionString) ImplementsObservabilityTelemetryLiveTailParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryLiveTailParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion() {}
func (UnionString) ImplementsObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityQueryNewResponseParametersNeedleValueUnion() {}
func (UnionString) ImplementsObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityQueryListResponseParametersNeedleValueUnion() {}
func (UnionString) ImplementsObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityQueryNewParamsParametersNeedleValueUnion() {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersNeedleValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion() {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseEventsEventsSourceUnion()           {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion()       {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseInvocationsSourceUnion()            {}
func (UnionString) ImplementsObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryNewParamsParametersNeedleValueUnion()         {}
func (UnionString) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()           {}
func (UnionString) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()        {}
func (UnionString) ImplementsNamespaceValueUpdateParamsValueUnion()                                {}
func (UnionString) ImplementsOriginPortUnionParam()                                                {}
func (UnionString) ImplementsOriginPortUnion()                                                     {}
func (UnionString) ImplementsV2QueryGetResponseEnvelopeErrorsCode()                                {}
func (UnionString) ImplementsV2QueryGetResponseEnvelopeMessagesCode()                              {}
func (UnionString) ImplementsHealthCheckTargetUnionParam()                                         {}
func (UnionString) ImplementsHealthCheckTargetUnion()                                              {}
func (UnionString) ImplementsGRETunnelNewResponseHealthCheckTargetUnion()                          {}
func (UnionString) ImplementsGRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion()      {}
func (UnionString) ImplementsGRETunnelListResponseGRETunnelsHealthCheckTargetUnion()               {}
func (UnionString) ImplementsGRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion()       {}
func (UnionString) ImplementsGRETunnelBulkUpdateResponseModifiedGRETunnelsHealthCheckTargetUnion() {}
func (UnionString) ImplementsGRETunnelGetResponseGRETunnelHealthCheckTargetUnion()                 {}
func (UnionString) ImplementsGRETunnelNewParamsHealthCheckTargetUnion()                            {}
func (UnionString) ImplementsGRETunnelUpdateParamsHealthCheckTargetUnion()                         {}
func (UnionString) ImplementsIPSECTunnelNewResponseHealthCheckTargetUnion()                        {}
func (UnionString) ImplementsIPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion()  {}
func (UnionString) ImplementsIPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion()           {}
func (UnionString) ImplementsIPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion()   {}
func (UnionString) ImplementsIPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion()         {}
func (UnionString) ImplementsIPSECTunnelNewParamsHealthCheckTargetUnion()                      {}
func (UnionString) ImplementsIPSECTunnelUpdateParamsHealthCheckTargetUnion()                   {}
func (UnionString) ImplementsAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodySaaSApplicationPolicyUnion()        {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationNewParamsBodyGatewayIdentityProxyEndpointApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBookmarkApplicationPolicyUnion()        {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserRDPApplicationPolicyUnion()      {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyMcpServerApplicationPolicyUnion()       {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyMcpServerPortalApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion()   {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion()         {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion()   {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion()   {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyGatewayIdentityProxyEndpointApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBookmarkApplicationPolicyUnion()   {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserRDPApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyMcpServerApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyMcpServerPortalApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationPolicyTestNewParamsPolicyUnion()                  {}
func (UnionString) ImplementsCasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion() {}
func (UnionString) ImplementsCasbPostureRemediationJobListResponseAssetFieldsValueUnion()       {}
func (UnionString) ImplementsDLPEmailRuleNewResponseConditionsValueUnion()                      {}
func (UnionString) ImplementsDLPEmailRuleUpdateResponseConditionsValueUnion()                   {}
func (UnionString) ImplementsDLPEmailRuleListResponseConditionsValueUnion()                     {}
func (UnionString) ImplementsDLPEmailRuleDeleteResponseConditionsValueUnion()                   {}
func (UnionString) ImplementsDLPEmailRuleBulkEditResponseConditionsValueUnion()                 {}
func (UnionString) ImplementsDLPEmailRuleGetResponseConditionsValueUnion()                      {}
func (UnionString) ImplementsDLPEmailRuleNewParamsConditionsValueUnion()                        {}
func (UnionString) ImplementsDLPEmailRuleUpdateParamsConditionsValueUnion()                     {}
func (UnionString) ImplementsRankingTimeseriesGroupsResponseSerie0Union()                       {}
func (UnionString) ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()        {}
func (UnionString) ImplementsConfigurationToolsZarazManagedComponentDefaultFieldsUnion()        {}
func (UnionString) ImplementsConfigurationToolsZarazManagedComponentSettingsUnion()             {}
func (UnionString) ImplementsConfigurationToolsWorkerDefaultFieldsUnion()                       {}
func (UnionString) ImplementsConfigurationToolsWorkerSettingsUnion()                            {}
func (UnionString) ImplementsConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion()   {}
func (UnionString) ImplementsConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()        {}
func (UnionString) ImplementsConfigUpdateParamsToolsWorkerDefaultFieldsUnion()                  {}
func (UnionString) ImplementsConfigUpdateParamsToolsWorkerSettingsUnion()                       {}
func (UnionString) ImplementsSessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion() {
}
func (UnionString) ImplementsThreatEventListParamsSearchValueUnion()                          {}
func (UnionString) ImplementsThreatEventListParamsSearchValueArrayItemUnion()                 {}
func (UnionString) ImplementsThreatEventRelationshipListParamsRelationshipTypesUnion()        {}
func (UnionString) ImplementsThreatEventIndicatorListParamsSearchValueUnion()                 {}
func (UnionString) ImplementsThreatEventIndicatorListParamsTagSearchValueUnion()              {}
func (UnionString) ImplementsThreatEventIndicatorListParamsTagSearchValueArrayItemUnion()     {}
func (UnionString) ImplementsThreatEventTagListParamsFiltersValueUnion()                      {}
func (UnionString) ImplementsThreatEventTagListParamsFiltersValueArrayItemUnion()             {}
func (UnionString) ImplementsThreatEventTagIndicatorListParamsSearchValueUnion()              {}
func (UnionString) ImplementsThreatEventTagIndicatorByDatasetListParamsSearchValueUnion()     {}
func (UnionString) ImplementsLogListParamsFiltersValueUnion()                                 {}
func (UnionString) ImplementsLogDeleteParamsFiltersValueUnion()                               {}
func (UnionString) ImplementsLogEditParamsMetadataUnion()                                     {}
func (UnionString) ImplementsDatasetNewResponseFiltersValueUnion()                            {}
func (UnionString) ImplementsDatasetUpdateResponseFiltersValueUnion()                         {}
func (UnionString) ImplementsDatasetListResponseFiltersValueUnion()                           {}
func (UnionString) ImplementsDatasetDeleteResponseFiltersValueUnion()                         {}
func (UnionString) ImplementsDatasetGetResponseFiltersValueUnion()                            {}
func (UnionString) ImplementsDatasetNewParamsFiltersValueUnion()                              {}
func (UnionString) ImplementsDatasetUpdateParamsFiltersValueUnion()                           {}
func (UnionString) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                 {}
func (UnionString) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                {}
func (UnionString) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()              {}
func (UnionString) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                 {}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseVariationsUnion()                                 {}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseVariationsUnion()                            {}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionString) ImplementsAppFlagListResponseVariationsUnion()                             {}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseVariationsUnion()                            {}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsVariationsUnion()                                 {}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectValueUnion()              {}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsVariationsUnion() {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()              {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                     {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                       {}
func (UnionString) ImplementsAppEvaluateGetResponseValueUnion()                                    {}
func (UnionString) ImplementsWorkflowUpdateParamsDefaultRetentionErrorRetentionUnion()             {}
func (UnionString) ImplementsWorkflowUpdateParamsDefaultRetentionSuccessRetentionUnion()           {}
func (UnionString) ImplementsInstanceGetResponseOutputUnion()                                      {}
func (UnionString) ImplementsInstanceGetResponseStepsObjectConfigRetriesDelayUnion()               {}
func (UnionString) ImplementsInstanceGetResponseStepsObjectConfigTimeoutUnion()                    {}
func (UnionString) ImplementsInstanceNewParamsInstanceRetentionErrorRetentionUnion()               {}
func (UnionString) ImplementsInstanceNewParamsInstanceRetentionSuccessRetentionUnion()             {}
func (UnionString) ImplementsInstanceBulkParamsBodyInstanceRetentionErrorRetentionUnion()          {}
func (UnionString) ImplementsInstanceBulkParamsBodyInstanceRetentionSuccessRetentionUnion()        {}
func (UnionString) ImplementsVersionGraphResponseGraphWorkflowNodesObjectDurationUnion()           {}
func (UnionString) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigRetriesDelayUnion() {}
func (UnionString) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigTimeoutUnion()      {}
func (UnionString) ImplementsVersionGraphResponseGraphWorkflowNodesObjectOptionsTimeoutUnion()     {}
func (UnionString) ImplementsAbuseReportGetResponseEnvelopeErrorsCode()                            {}
func (UnionString) ImplementsAIRunResponseUnion()                                                  {}
func (UnionString) ImplementsAIRunParamsBodyTextEmbeddingsTextUnion()                              {}
func (UnionString) ImplementsAIRunParamsBodyTextGenerationMessagesContentUnion()                   {}
func (UnionString) ImplementsAIRunParamsBodyImageTextToTextMessagesContentUnion()                  {}
func (UnionString) ImplementsNamespaceChatCompletionsResponseChoicesMessageContentUnion()          {}
func (UnionString) ImplementsNamespaceChatCompletionsParamsMessagesContentUnion()                  {}
func (UnionString) ImplementsNamespaceSearchParamsMessagesContentUnion()                           {}
func (UnionString) ImplementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion()  {}
func (UnionString) ImplementsNamespaceInstanceChatCompletionsParamsMessagesContentUnion()          {}
func (UnionString) ImplementsNamespaceInstanceSearchParamsMessagesContentUnion()                   {}
func (UnionString) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                      {}
func (UnionString) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()               {}
func (UnionString) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                       {}
func (UnionString) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                      {}
func (UnionString) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                    {}
func (UnionString) ImplementsInstanceChatCompletionsResponseChoicesMessageContentUnion()           {}
func (UnionString) ImplementsInstanceChatCompletionsParamsMessagesContentUnion()                   {}
func (UnionString) ImplementsInstanceSearchParamsMessagesContentUnion()                            {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsHeightUnion()                                   {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginBottomUnion()                             {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginLeftUnion()                               {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginRightUnion()                              {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginTopUnion()                                {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsWidthUnion()                                    {}
func (UnionString) ImplementsSnapshotNewResponseAccessibilityTreeValueUnion()                      {}
func (UnionString) ImplementsJsonNewParamsResponseFormatJsonSchemaUnion()                          {}
func (UnionString) ImplementsAccessibilityTreeNewResponseAccessibilityTreeValueUnion()             {}
func (UnionString) ImplementsCrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion()    {}

type UnionBool bool

func (UnionBool) ImplementsAnalyticsQuerySummaryParamsFiltersValueUnion()                        {}
func (UnionBool) ImplementsAnalyticsQueryTimeseriesParamsFiltersValueUnion()                     {}
func (UnionBool) ImplementsAnalyticsQueryTopNParamsFiltersValueUnion()                           {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion() {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion()     {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion()  {}
func (UnionBool) ImplementsVersionAssetsConfigRunWorkerFirstUnionParam()                         {}
func (UnionBool) ImplementsVersionAssetsConfigRunWorkerFirstUnion()                              {}
func (UnionBool) ImplementsScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()           {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()       {}
func (UnionBool) ImplementsObservabilityTelemetryValuesResponseValueUnion()                       {}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()    {}
func (UnionBool) ImplementsObservabilityTelemetryLiveTailParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryLiveTailParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityQueryNewResponseParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityQueryListResponseParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilityQueryNewParamsParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersNeedleValueUnion() {}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion() {}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionBool) ImplementsObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion()       {}
func (UnionBool) ImplementsObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionBool) ImplementsObservabilitySharedQueryNewParamsParametersNeedleValueUnion()  {}
func (UnionBool) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()    {}
func (UnionBool) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion() {}
func (UnionBool) ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}
func (UnionBool) ImplementsCasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion() {}
func (UnionBool) ImplementsCasbPostureRemediationJobListResponseAssetFieldsValueUnion()       {}
func (UnionBool) ImplementsConfigurationToolsZarazManagedComponentDefaultFieldsUnion()        {}
func (UnionBool) ImplementsConfigurationToolsZarazManagedComponentSettingsUnion()             {}
func (UnionBool) ImplementsConfigurationToolsWorkerDefaultFieldsUnion()                       {}
func (UnionBool) ImplementsConfigurationToolsWorkerSettingsUnion()                            {}
func (UnionBool) ImplementsConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion()   {}
func (UnionBool) ImplementsConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()        {}
func (UnionBool) ImplementsConfigUpdateParamsToolsWorkerDefaultFieldsUnion()                  {}
func (UnionBool) ImplementsConfigUpdateParamsToolsWorkerSettingsUnion()                       {}
func (UnionBool) ImplementsSessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion() {
}
func (UnionBool) ImplementsLogListParamsFiltersValueUnion()                                 {}
func (UnionBool) ImplementsLogDeleteParamsFiltersValueUnion()                               {}
func (UnionBool) ImplementsLogEditParamsMetadataUnion()                                     {}
func (UnionBool) ImplementsDatasetNewResponseFiltersValueUnion()                            {}
func (UnionBool) ImplementsDatasetUpdateResponseFiltersValueUnion()                         {}
func (UnionBool) ImplementsDatasetListResponseFiltersValueUnion()                           {}
func (UnionBool) ImplementsDatasetDeleteResponseFiltersValueUnion()                         {}
func (UnionBool) ImplementsDatasetGetResponseFiltersValueUnion()                            {}
func (UnionBool) ImplementsDatasetNewParamsFiltersValueUnion()                              {}
func (UnionBool) ImplementsDatasetUpdateParamsFiltersValueUnion()                           {}
func (UnionBool) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                 {}
func (UnionBool) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                {}
func (UnionBool) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()              {}
func (UnionBool) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                 {}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseVariationsUnion()                                 {}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseVariationsUnion()                            {}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseVariationsUnion()                             {}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseVariationsUnion()                            {}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsVariationsUnion()                                 {}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectValueUnion()              {}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsVariationsUnion()                                     {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectValueUnion() {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()           {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                  {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                    {}
func (UnionBool) ImplementsAppEvaluateGetResponseValueUnion()                                 {}
func (UnionBool) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                   {}
func (UnionBool) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()            {}
func (UnionBool) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                    {}
func (UnionBool) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                   {}
func (UnionBool) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                 {}
func (UnionBool) ImplementsSnapshotNewResponseAccessibilityTreeCheckedUnion()                 {}
func (UnionBool) ImplementsSnapshotNewResponseAccessibilityTreePressedUnion()                 {}
func (UnionBool) ImplementsJsonNewParamsResponseFormatJsonSchemaUnion()                       {}
func (UnionBool) ImplementsAccessibilityTreeNewResponseAccessibilityTreeCheckedUnion()        {}
func (UnionBool) ImplementsAccessibilityTreeNewResponseAccessibilityTreePressedUnion()        {}
func (UnionBool) ImplementsCrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion() {}

type UnionInt int64

func (UnionInt) ImplementsReceivedGetParamsEndUnion()                                    {}
func (UnionInt) ImplementsReceivedGetParamsStartUnion()                                  {}
func (UnionInt) ImplementsPageRuleActionsCacheTTLByStatusValueUnion()                    {}
func (UnionInt) ImplementsPageRuleNewParamsActionsCacheTTLByStatusValueUnion()           {}
func (UnionInt) ImplementsPageRuleUpdateParamsActionsCacheTTLByStatusValueUnion()        {}
func (UnionInt) ImplementsPageRuleEditParamsActionsCacheTTLByStatusValueUnion()          {}
func (UnionInt) ImplementsOriginPortUnionParam()                                         {}
func (UnionInt) ImplementsOriginPortUnion()                                              {}
func (UnionInt) ImplementsWorkflowUpdateParamsDefaultRetentionErrorRetentionUnion()      {}
func (UnionInt) ImplementsWorkflowUpdateParamsDefaultRetentionSuccessRetentionUnion()    {}
func (UnionInt) ImplementsInstanceNewParamsInstanceRetentionErrorRetentionUnion()        {}
func (UnionInt) ImplementsInstanceNewParamsInstanceRetentionSuccessRetentionUnion()      {}
func (UnionInt) ImplementsInstanceBulkParamsBodyInstanceRetentionErrorRetentionUnion()   {}
func (UnionInt) ImplementsInstanceBulkParamsBodyInstanceRetentionSuccessRetentionUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsSettingEditParamsBodyValueValueUnion()                                {}
func (UnionFloat) ImplementsAnalyticsQuerySummaryParamsFiltersValueUnion()                        {}
func (UnionFloat) ImplementsAnalyticsQueryTimeseriesParamsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsAnalyticsQueryTopNParamsFiltersValueUnion()                           {}
func (UnionFloat) ImplementsAnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion() {}
func (UnionFloat) ImplementsAnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion()     {}
func (UnionFloat) ImplementsAnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion()  {}
func (UnionFloat) ImplementsTTLParam()                                                            {}
func (UnionFloat) ImplementsTTL()                                                                 {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()       {}
func (UnionFloat) ImplementsObservabilityTelemetryValuesResponseValueUnion()                       {}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()    {}
func (UnionFloat) ImplementsObservabilityTelemetryLiveTailParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryLiveTailParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityQueryNewResponseParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityQueryListResponseParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilityQueryNewParamsParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion()       {}
func (UnionFloat) ImplementsObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionFloat) ImplementsObservabilitySharedQueryNewParamsParametersNeedleValueUnion()      {}
func (UnionFloat) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()        {}
func (UnionFloat) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()     {}
func (UnionFloat) ImplementsV2QueryGetResponseEnvelopeErrorsCode()                             {}
func (UnionFloat) ImplementsV2QueryGetResponseEnvelopeMessagesCode()                           {}
func (UnionFloat) ImplementsCasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion() {}
func (UnionFloat) ImplementsCasbPostureRemediationJobListResponseAssetFieldsValueUnion()       {}
func (UnionFloat) ImplementsRankingTimeseriesGroupsResponseSerie0Union()                       {}
func (UnionFloat) ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()        {}
func (UnionFloat) ImplementsSessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion() {
}
func (UnionFloat) ImplementsThreatEventListParamsSearchValueUnion()                          {}
func (UnionFloat) ImplementsThreatEventListParamsSearchValueArrayItemUnion()                 {}
func (UnionFloat) ImplementsThreatEventIndicatorListParamsTagSearchValueUnion()              {}
func (UnionFloat) ImplementsThreatEventIndicatorListParamsTagSearchValueArrayItemUnion()     {}
func (UnionFloat) ImplementsThreatEventTagListParamsFiltersValueUnion()                      {}
func (UnionFloat) ImplementsThreatEventTagListParamsFiltersValueArrayItemUnion()             {}
func (UnionFloat) ImplementsLogListParamsFiltersValueUnion()                                 {}
func (UnionFloat) ImplementsLogDeleteParamsFiltersValueUnion()                               {}
func (UnionFloat) ImplementsLogEditParamsMetadataUnion()                                     {}
func (UnionFloat) ImplementsDatasetNewResponseFiltersValueUnion()                            {}
func (UnionFloat) ImplementsDatasetUpdateResponseFiltersValueUnion()                         {}
func (UnionFloat) ImplementsDatasetListResponseFiltersValueUnion()                           {}
func (UnionFloat) ImplementsDatasetDeleteResponseFiltersValueUnion()                         {}
func (UnionFloat) ImplementsDatasetGetResponseFiltersValueUnion()                            {}
func (UnionFloat) ImplementsDatasetNewParamsFiltersValueUnion()                              {}
func (UnionFloat) ImplementsDatasetUpdateParamsFiltersValueUnion()                           {}
func (UnionFloat) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                 {}
func (UnionFloat) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                {}
func (UnionFloat) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()              {}
func (UnionFloat) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                 {}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseVariationsUnion()                                 {}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseVariationsUnion()                            {}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseVariationsUnion()                             {}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseVariationsUnion()                            {}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsVariationsUnion()                                 {}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectValueUnion()              {}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectValueUnion() {}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClauseUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsVariationsUnion() {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesObjectClausesUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()              {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                     {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                       {}
func (UnionFloat) ImplementsAppEvaluateGetResponseValueUnion()                                    {}
func (UnionFloat) ImplementsInstanceGetResponseOutputUnion()                                      {}
func (UnionFloat) ImplementsInstanceGetResponseStepsObjectConfigRetriesDelayUnion()               {}
func (UnionFloat) ImplementsInstanceGetResponseStepsObjectConfigTimeoutUnion()                    {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectDurationUnion()           {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigRetriesDelayUnion() {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigTimeoutUnion()      {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectOptionsTimeoutUnion()     {}
func (UnionFloat) ImplementsAbuseReportGetResponseEnvelopeErrorsCode()                            {}
func (UnionFloat) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                      {}
func (UnionFloat) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()               {}
func (UnionFloat) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                       {}
func (UnionFloat) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                      {}
func (UnionFloat) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                    {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsHeightUnion()                                   {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginBottomUnion()                             {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginLeftUnion()                               {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginRightUnion()                              {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginTopUnion()                                {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsWidthUnion()                                    {}
func (UnionFloat) ImplementsSnapshotNewResponseAccessibilityTreeValueUnion()                      {}
func (UnionFloat) ImplementsJsonNewParamsResponseFormatJsonSchemaUnion()                          {}
func (UnionFloat) ImplementsAccessibilityTreeNewResponseAccessibilityTreeValueUnion()             {}
func (UnionFloat) ImplementsCrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion()    {}

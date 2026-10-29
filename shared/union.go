// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsAuditLogListParamsBeforeUnion()                             {}
func (UnionTime) ImplementsAuditLogListParamsSinceUnion()                              {}
func (UnionTime) ImplementsObservabilityTelemetryQueryResponseRunQueryCreatedUnion()   {}
func (UnionTime) ImplementsObservabilityTelemetryQueryResponseRunQueryUpdatedUnion()   {}
func (UnionTime) ImplementsObservabilitySharedQueryGetResponseRunQueryCreatedUnion()   {}
func (UnionTime) ImplementsObservabilitySharedQueryGetResponseRunQueryUpdatedUnion()   {}
func (UnionTime) ImplementsThreatEventIndicatorAggregateListParamsCreatedAfterUnion()  {}
func (UnionTime) ImplementsThreatEventIndicatorAggregateListParamsCreatedBeforeUnion() {}

type UnionString string

func (UnionString) ImplementsEntitlementListResponseAllocationValueUnion()                         {}
func (UnionString) ImplementsBillingUsageGetResponseTagsUnion()                                    {}
func (UnionString) ImplementsSettingEditParamsBodyValueValueUnion()                                {}
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
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryCreatedUnion()             {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryUpdatedUnion() {}
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
func (UnionString) ImplementsObservabilityQueryNewParamsParametersNeedleValueUnion()   {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryCreatedUnion() {}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryParametersNeedleValueUnion() {
}
func (UnionString) ImplementsObservabilitySharedQueryGetResponseRunQueryUpdatedUnion() {}
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
func (UnionString) ImplementsObservabilitySharedQueryNewParamsParametersNeedleValueUnion()  {}
func (UnionString) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()    {}
func (UnionString) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion() {}
func (UnionString) ImplementsNamespaceValueUpdateParamsValueUnion()                         {}
func (UnionString) ImplementsRulesetNewResponseUnion()                                      {}
func (UnionString) ImplementsRulesetUpdateResponseUnion()                                   {}
func (UnionString) ImplementsRulesetListResponseUnion()                                     {}
func (UnionString) ImplementsPhaseUpdateResponseUnion()                                     {}
func (UnionString) ImplementsPhaseVersionListResponseUnion()                                {}
func (UnionString) ImplementsRuleNewResponseUnion()                                         {}
func (UnionString) ImplementsRuleDeleteResponseUnion()                                      {}
func (UnionString) ImplementsRuleEditResponseUnion()                                        {}
func (UnionString) ImplementsVersionListResponseUnion()                                     {}
func (UnionString) ImplementsOriginPortUnionParam()                                         {}
func (UnionString) ImplementsOriginPortUnion()                                              {}
func (UnionString) ImplementsUsageGetResponseTagsUnion()                                    {}
func (UnionString) ImplementsUsageGetAccountUsageV2ResponseTagsUnion()                      {}
func (UnionString) ImplementsV2QueryGetResponseEnvelopeErrorsCode()                         {}
func (UnionString) ImplementsV2QueryGetResponseEnvelopeMessagesCode()                       {}
func (UnionString) ImplementsCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion() {
}
func (UnionString) ImplementsCfInterconnectListResponseInterconnectsHealthCheckTargetUnion() {}
func (UnionString) ImplementsCfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsCfInterconnectGetResponseInterconnectHealthCheckTargetUnion()         {}
func (UnionString) ImplementsCfInterconnectUpdateParamsHealthCheckTargetUnion()                    {}
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
func (UnionString) ImplementsThreatEventListParamsSearchObjectValueUnion()                         {}
func (UnionString) ImplementsThreatEventListParamsSearchBranchesObjectValueUnion()                 {}
func (UnionString) ImplementsThreatEventRelationshipListParamsRelationshipTypesUnion()             {}
func (UnionString) ImplementsThreatEventIndicatorListParamsSearchValueUnion()                      {}
func (UnionString) ImplementsThreatEventIndicatorListParamsTagSearchValueUnion()                   {}
func (UnionString) ImplementsThreatEventIndicatorListParamsTagSearchValueArrayItemUnion()          {}
func (UnionString) ImplementsThreatEventIndicatorAggregateListParamsCreatedAfterUnion()            {}
func (UnionString) ImplementsThreatEventIndicatorAggregateListParamsCreatedBeforeUnion()           {}
func (UnionString) ImplementsThreatEventTagNewParamsActiveDurationUnion()                          {}
func (UnionString) ImplementsThreatEventTagNewParamsActorCategoryUnion()                           {}
func (UnionString) ImplementsThreatEventTagNewParamsAttributionOrganizationUnion()                 {}
func (UnionString) ImplementsThreatEventTagNewParamsMotiveUnion()                                  {}
func (UnionString) ImplementsThreatEventTagNewParamsOpsecLevelUnion()                              {}
func (UnionString) ImplementsThreatEventTagNewParamsOriginCountryISOUnion()                        {}
func (UnionString) ImplementsThreatEventTagNewParamsSophisticationLevelUnion()                     {}
func (UnionString) ImplementsThreatEventTagListParamsFiltersValueUnion()                           {}
func (UnionString) ImplementsThreatEventTagListParamsFiltersValueArrayItemUnion()                  {}
func (UnionString) ImplementsThreatEventTagEditParamsActiveDurationUnion()                         {}
func (UnionString) ImplementsThreatEventTagEditParamsActorCategoryUnion()                          {}
func (UnionString) ImplementsThreatEventTagEditParamsAttributionOrganizationUnion()                {}
func (UnionString) ImplementsThreatEventTagEditParamsMotiveUnion()                                 {}
func (UnionString) ImplementsThreatEventTagEditParamsOpsecLevelUnion()                             {}
func (UnionString) ImplementsThreatEventTagEditParamsOriginCountryISOUnion()                       {}
func (UnionString) ImplementsThreatEventTagEditParamsSophisticationLevelUnion()                    {}
func (UnionString) ImplementsThreatEventTagIndicatorListParamsSearchValueUnion()                   {}
func (UnionString) ImplementsThreatEventTagIndicatorByDatasetListParamsSearchValueUnion()          {}
func (UnionString) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion()          {}
func (UnionString) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem() {}
func (UnionString) ImplementsLogListParamsFiltersValueUnion()                                      {}
func (UnionString) ImplementsLogDeleteParamsFiltersValueUnion()                                    {}
func (UnionString) ImplementsLogEditParamsMetadataUnion()                                          {}
func (UnionString) ImplementsDatasetNewResponseFiltersValueUnion()                                 {}
func (UnionString) ImplementsDatasetUpdateResponseFiltersValueUnion()                              {}
func (UnionString) ImplementsDatasetListResponseFiltersValueUnion()                                {}
func (UnionString) ImplementsDatasetDeleteResponseFiltersValueUnion()                              {}
func (UnionString) ImplementsDatasetGetResponseFiltersValueUnion()                                 {}
func (UnionString) ImplementsDatasetNewParamsFiltersValueUnion()                                   {}
func (UnionString) ImplementsDatasetUpdateParamsFiltersValueUnion()                                {}
func (UnionString) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                      {}
func (UnionString) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                     {}
func (UnionString) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()                   {}
func (UnionString) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                      {}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion()    {}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagNewResponseVariationsUnion()                                   {}
func (UnionString) ImplementsAppFlagNewResponseMetadataUnion()                                     {}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateResponseVariationsUnion()                              {}
func (UnionString) ImplementsAppFlagUpdateResponseMetadataUnion()                                {}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagListResponseVariationsUnion()                               {}
func (UnionString) ImplementsAppFlagListResponseMetadataUnion()                                 {}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagGetResponseVariationsUnion()                              {}
func (UnionString) ImplementsAppFlagGetResponseMetadataUnion()                                {}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagNewParamsVariationsUnion()                                   {}
func (UnionString) ImplementsAppFlagNewParamsMetadataUnion()                                     {}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagUpdateParamsVariationsUnion() {}
func (UnionString) ImplementsAppFlagUpdateParamsMetadataUnion()   {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()              {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectAfterMetadataUnion()                {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                     {}
func (UnionString) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                       {}
func (UnionString) ImplementsAppEvaluateGetResponseMetadataUnion()                                 {}
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
func (UnionString) ImplementsSubmittedListResponseEnvelopeErrorsCode()                             {}
func (UnionString) ImplementsSubmittedGetResponseEnvelopeErrorsCode()                              {}
func (UnionString) ImplementsSubmittedEmailListResponseEnvelopeErrorsCode()                        {}
func (UnionString) ImplementsAIRunResponseUnion()                                                  {}
func (UnionString) ImplementsAIRunParamsBodyTextEmbeddingsTextUnion()                              {}
func (UnionString) ImplementsAIRunParamsBodyTextGenerationMessagesContentUnion()                   {}
func (UnionString) ImplementsAIRunParamsBodyImageTextToTextMessagesContentUnion()                  {}
func (UnionString) ImplementsNamespaceChatCompletionsResponseChoicesMessageContentUnion()          {}
func (UnionString) ImplementsNamespaceChatCompletionsParamsAISearchOptionsCustomMetadataUnion()    {}
func (UnionString) ImplementsNamespaceChatCompletionsParamsMessagesContentUnion()                  {}
func (UnionString) ImplementsNamespaceSearchParamsAISearchOptionsCustomMetadataUnion()             {}
func (UnionString) ImplementsNamespaceSearchParamsMessagesContentUnion()                           {}
func (UnionString) ImplementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion()  {}
func (UnionString) ImplementsNamespaceInstanceChatCompletionsParamsMessagesContentUnion()          {}
func (UnionString) ImplementsNamespaceInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {
}
func (UnionString) ImplementsNamespaceInstanceSearchParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionString) ImplementsNamespaceInstanceSearchParamsMessagesContentUnion()               {}
func (UnionString) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                  {}
func (UnionString) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()           {}
func (UnionString) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                   {}
func (UnionString) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                  {}
func (UnionString) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                {}
func (UnionString) ImplementsInstanceChatCompletionsResponseChoicesMessageContentUnion()       {}
func (UnionString) ImplementsInstanceChatCompletionsParamsMessagesContentUnion()               {}
func (UnionString) ImplementsInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionString) ImplementsInstanceSearchParamsAISearchOptionsCustomMetadataUnion()          {}
func (UnionString) ImplementsInstanceSearchParamsMessagesContentUnion()                        {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsHeightUnion()                     {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginBottomUnion()               {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginLeftUnion()                 {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginRightUnion()                {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginTopUnion()                  {}
func (UnionString) ImplementsPDFNewParamsBodyObjectPDFOptionsWidthUnion()                      {}
func (UnionString) ImplementsSnapshotNewResponseAccessibilityTreeValueUnion()                  {}
func (UnionString) ImplementsAccessibilityTreeNewResponseAccessibilityTreeValueUnion()         {}

type UnionBool bool

func (UnionBool) ImplementsEntitlementListResponseAllocationValueUnion()                         {}
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
func (UnionBool) ImplementsRulesetNewResponseUnion()                                      {}
func (UnionBool) ImplementsRulesetUpdateResponseUnion()                                   {}
func (UnionBool) ImplementsRulesetListResponseUnion()                                     {}
func (UnionBool) ImplementsPhaseUpdateResponseUnion()                                     {}
func (UnionBool) ImplementsPhaseVersionListResponseUnion()                                {}
func (UnionBool) ImplementsRuleNewResponseUnion()                                         {}
func (UnionBool) ImplementsRuleDeleteResponseUnion()                                      {}
func (UnionBool) ImplementsRuleEditResponseUnion()                                        {}
func (UnionBool) ImplementsVersionListResponseUnion()                                     {}
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
func (UnionBool) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion()          {}
func (UnionBool) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem() {}
func (UnionBool) ImplementsLogListParamsFiltersValueUnion()                                      {}
func (UnionBool) ImplementsLogDeleteParamsFiltersValueUnion()                                    {}
func (UnionBool) ImplementsLogEditParamsMetadataUnion()                                          {}
func (UnionBool) ImplementsDatasetNewResponseFiltersValueUnion()                                 {}
func (UnionBool) ImplementsDatasetUpdateResponseFiltersValueUnion()                              {}
func (UnionBool) ImplementsDatasetListResponseFiltersValueUnion()                                {}
func (UnionBool) ImplementsDatasetDeleteResponseFiltersValueUnion()                              {}
func (UnionBool) ImplementsDatasetGetResponseFiltersValueUnion()                                 {}
func (UnionBool) ImplementsDatasetNewParamsFiltersValueUnion()                                   {}
func (UnionBool) ImplementsDatasetUpdateParamsFiltersValueUnion()                                {}
func (UnionBool) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                      {}
func (UnionBool) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                     {}
func (UnionBool) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()                   {}
func (UnionBool) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                      {}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion()    {}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewResponseVariationsUnion()                                   {}
func (UnionBool) ImplementsAppFlagNewResponseMetadataUnion()                                     {}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateResponseVariationsUnion()                              {}
func (UnionBool) ImplementsAppFlagUpdateResponseMetadataUnion()                                {}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagListResponseVariationsUnion()                               {}
func (UnionBool) ImplementsAppFlagListResponseMetadataUnion()                                 {}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagGetResponseVariationsUnion()                              {}
func (UnionBool) ImplementsAppFlagGetResponseMetadataUnion()                                {}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagNewParamsVariationsUnion()                                   {}
func (UnionBool) ImplementsAppFlagNewParamsMetadataUnion()                                     {}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagUpdateParamsVariationsUnion() {}
func (UnionBool) ImplementsAppFlagUpdateParamsMetadataUnion()   {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()           {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectAfterMetadataUnion()             {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                  {}
func (UnionBool) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                    {}
func (UnionBool) ImplementsAppEvaluateGetResponseMetadataUnion()                              {}
func (UnionBool) ImplementsAppEvaluateGetResponseValueUnion()                                 {}
func (UnionBool) ImplementsNamespaceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionBool) ImplementsNamespaceSearchParamsAISearchOptionsCustomMetadataUnion()          {}
func (UnionBool) ImplementsNamespaceInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {
}
func (UnionBool) ImplementsNamespaceInstanceSearchParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionBool) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                  {}
func (UnionBool) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()           {}
func (UnionBool) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                   {}
func (UnionBool) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                  {}
func (UnionBool) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                {}
func (UnionBool) ImplementsInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionBool) ImplementsInstanceSearchParamsAISearchOptionsCustomMetadataUnion()          {}
func (UnionBool) ImplementsSnapshotNewResponseAccessibilityTreeCheckedUnion()                {}
func (UnionBool) ImplementsSnapshotNewResponseAccessibilityTreePressedUnion()                {}
func (UnionBool) ImplementsAccessibilityTreeNewResponseAccessibilityTreeCheckedUnion()       {}
func (UnionBool) ImplementsAccessibilityTreeNewResponseAccessibilityTreePressedUnion()       {}

type UnionInt int64

func (UnionInt) ImplementsEntitlementListResponseAllocationValueUnion()                  {}
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
func (UnionFloat) ImplementsRulesetNewResponseUnion()                                          {}
func (UnionFloat) ImplementsRulesetUpdateResponseUnion()                                       {}
func (UnionFloat) ImplementsRulesetListResponseUnion()                                         {}
func (UnionFloat) ImplementsPhaseUpdateResponseUnion()                                         {}
func (UnionFloat) ImplementsPhaseVersionListResponseUnion()                                    {}
func (UnionFloat) ImplementsRuleNewResponseUnion()                                             {}
func (UnionFloat) ImplementsRuleDeleteResponseUnion()                                          {}
func (UnionFloat) ImplementsRuleEditResponseUnion()                                            {}
func (UnionFloat) ImplementsVersionListResponseUnion()                                         {}
func (UnionFloat) ImplementsV2QueryGetResponseEnvelopeErrorsCode()                             {}
func (UnionFloat) ImplementsV2QueryGetResponseEnvelopeMessagesCode()                           {}
func (UnionFloat) ImplementsCasbPostureRemediationJobNewResponseCreatedAssetFieldsValueUnion() {}
func (UnionFloat) ImplementsCasbPostureRemediationJobListResponseAssetFieldsValueUnion()       {}
func (UnionFloat) ImplementsValidatorWorkerValidatorDefaultUnion()                             {}
func (UnionFloat) ImplementsRankingTimeseriesGroupsResponseSerie0Union()                       {}
func (UnionFloat) ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()        {}
func (UnionFloat) ImplementsSessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion() {
}
func (UnionFloat) ImplementsThreatEventListParamsSearchObjectValueUnion()                         {}
func (UnionFloat) ImplementsThreatEventListParamsSearchBranchesObjectValueUnion()                 {}
func (UnionFloat) ImplementsThreatEventIndicatorListParamsTagSearchValueUnion()                   {}
func (UnionFloat) ImplementsThreatEventIndicatorListParamsTagSearchValueArrayItemUnion()          {}
func (UnionFloat) ImplementsThreatEventTagNewParamsPriorityUnion()                                {}
func (UnionFloat) ImplementsThreatEventTagListParamsFiltersValueUnion()                           {}
func (UnionFloat) ImplementsThreatEventTagListParamsFiltersValueArrayItemUnion()                  {}
func (UnionFloat) ImplementsThreatEventTagEditParamsPriorityUnion()                               {}
func (UnionFloat) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion()          {}
func (UnionFloat) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem() {}
func (UnionFloat) ImplementsLogListParamsFiltersValueUnion()                                      {}
func (UnionFloat) ImplementsLogDeleteParamsFiltersValueUnion()                                    {}
func (UnionFloat) ImplementsLogEditParamsMetadataUnion()                                          {}
func (UnionFloat) ImplementsDatasetNewResponseFiltersValueUnion()                                 {}
func (UnionFloat) ImplementsDatasetUpdateResponseFiltersValueUnion()                              {}
func (UnionFloat) ImplementsDatasetListResponseFiltersValueUnion()                                {}
func (UnionFloat) ImplementsDatasetDeleteResponseFiltersValueUnion()                              {}
func (UnionFloat) ImplementsDatasetGetResponseFiltersValueUnion()                                 {}
func (UnionFloat) ImplementsDatasetNewParamsFiltersValueUnion()                                   {}
func (UnionFloat) ImplementsDatasetUpdateParamsFiltersValueUnion()                                {}
func (UnionFloat) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                      {}
func (UnionFloat) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()                   {}
func (UnionFloat) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                      {}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion()    {}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewResponseVariationsUnion()                                   {}
func (UnionFloat) ImplementsAppFlagNewResponseMetadataUnion()                                     {}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateResponseVariationsUnion()                              {}
func (UnionFloat) ImplementsAppFlagUpdateResponseMetadataUnion()                                {}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagListResponseVariationsUnion()                               {}
func (UnionFloat) ImplementsAppFlagListResponseMetadataUnion()                                 {}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagGetResponseVariationsUnion()                              {}
func (UnionFloat) ImplementsAppFlagGetResponseMetadataUnion()                                {}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagNewParamsVariationsUnion()                                   {}
func (UnionFloat) ImplementsAppFlagNewParamsMetadataUnion()                                     {}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion() {}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagUpdateParamsVariationsUnion() {}
func (UnionFloat) ImplementsAppFlagUpdateParamsMetadataUnion()   {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()              {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectAfterMetadataUnion()                {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectDiffFromUnion()                     {}
func (UnionFloat) ImplementsAppFlagChangelogListResponseObjectDiffToUnion()                       {}
func (UnionFloat) ImplementsAppEvaluateGetResponseMetadataUnion()                                 {}
func (UnionFloat) ImplementsAppEvaluateGetResponseValueUnion()                                    {}
func (UnionFloat) ImplementsInstanceGetResponseOutputUnion()                                      {}
func (UnionFloat) ImplementsInstanceGetResponseStepsObjectConfigRetriesDelayUnion()               {}
func (UnionFloat) ImplementsInstanceGetResponseStepsObjectConfigTimeoutUnion()                    {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectDurationUnion()           {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigRetriesDelayUnion() {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectConfigTimeoutUnion()      {}
func (UnionFloat) ImplementsVersionGraphResponseGraphWorkflowNodesObjectOptionsTimeoutUnion()     {}
func (UnionFloat) ImplementsAbuseReportGetResponseEnvelopeErrorsCode()                            {}
func (UnionFloat) ImplementsSubmittedListResponseEnvelopeErrorsCode()                             {}
func (UnionFloat) ImplementsSubmittedGetResponseEnvelopeErrorsCode()                              {}
func (UnionFloat) ImplementsSubmittedEmailListResponseEnvelopeErrorsCode()                        {}
func (UnionFloat) ImplementsNamespaceChatCompletionsParamsAISearchOptionsCustomMetadataUnion()    {}
func (UnionFloat) ImplementsNamespaceSearchParamsAISearchOptionsCustomMetadataUnion()             {}
func (UnionFloat) ImplementsNamespaceInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {
}
func (UnionFloat) ImplementsNamespaceInstanceSearchParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionFloat) ImplementsNamespaceInstanceItemListResponseMetadataUnion()                  {}
func (UnionFloat) ImplementsNamespaceInstanceItemNewOrUpdateResponseMetadataUnion()           {}
func (UnionFloat) ImplementsNamespaceInstanceItemGetResponseMetadataUnion()                   {}
func (UnionFloat) ImplementsNamespaceInstanceItemSyncResponseMetadataUnion()                  {}
func (UnionFloat) ImplementsNamespaceInstanceItemUploadResponseMetadataUnion()                {}
func (UnionFloat) ImplementsInstanceChatCompletionsParamsAISearchOptionsCustomMetadataUnion() {}
func (UnionFloat) ImplementsInstanceSearchParamsAISearchOptionsCustomMetadataUnion()          {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsHeightUnion()                     {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginBottomUnion()               {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginLeftUnion()                 {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginRightUnion()                {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsMarginTopUnion()                  {}
func (UnionFloat) ImplementsPDFNewParamsBodyObjectPDFOptionsWidthUnion()                      {}
func (UnionFloat) ImplementsSnapshotNewResponseAccessibilityTreeValueUnion()                  {}
func (UnionFloat) ImplementsAccessibilityTreeNewResponseAccessibilityTreeValueUnion()         {}

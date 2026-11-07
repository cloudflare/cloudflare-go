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

func (UnionString) ImplementsEntitlementListResponseAllocationValueUnion() {}
func (UnionString) ImplementsBillingUsageGetResponseTagsUnion()            {}
func (UnionString) ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion() {
}
func (UnionString) ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion() {
}
func (UnionString) ImplementsAdvancedDDoSValueUnion()                                              {}
func (UnionString) ImplementsAlwaysOnlineValueUnion()                                              {}
func (UnionString) ImplementsBrotliValueUnion()                                                    {}
func (UnionString) ImplementsChallengeTTLValueUnion()                                              {}
func (UnionString) ImplementsCiphersValueUnion()                                                   {}
func (UnionString) ImplementsDevelopmentModeValueUnion()                                           {}
func (UnionString) ImplementsEarlyHintsValueUnion()                                                {}
func (UnionString) ImplementsH2PrioritizationValueUnion()                                          {}
func (UnionString) ImplementsHotlinkProtectionValueUnion()                                         {}
func (UnionString) ImplementsHTTP2ValueUnion()                                                     {}
func (UnionString) ImplementsHTTP3ValueUnion()                                                     {}
func (UnionString) ImplementsImageResizingValueUnion()                                             {}
func (UnionString) ImplementsIPV6ValueUnion()                                                      {}
func (UnionString) ImplementsMinTLSVersionValueUnion()                                             {}
func (UnionString) ImplementsOpportunisticOnionValueUnion()                                        {}
func (UnionString) ImplementsOrangeToOrangeValueUnion()                                            {}
func (UnionString) ImplementsPrefetchPreloadValueUnion()                                           {}
func (UnionString) ImplementsProxyReadTimeoutValueUnion()                                          {}
func (UnionString) ImplementsPseudoIPV4ValueUnion()                                                {}
func (UnionString) ImplementsServerSideExcludesValueUnion()                                        {}
func (UnionString) ImplementsTls1_3ValueUnion()                                                    {}
func (UnionString) ImplementsTLSClientAuthValueUnion()                                             {}
func (UnionString) ImplementsWebPValueUnion()                                                      {}
func (UnionString) ImplementsWebsocketValueUnion()                                                 {}
func (UnionString) ImplementsZeroRTTValueUnion()                                                   {}
func (UnionString) ImplementsSettingEditResponseZonesAlwaysUseHttps2ValueUnion()                   {}
func (UnionString) ImplementsSettingEditResponseZonesAutomaticHTTPSRewrites2ValueUnion()           {}
func (UnionString) ImplementsSettingEditResponseZonesBrowserCacheTtl2ValueUnion()                  {}
func (UnionString) ImplementsSettingEditResponseZonesBrowserCheck2ValueUnion()                     {}
func (UnionString) ImplementsSettingEditResponseZonesCacheLevel2ValueUnion()                       {}
func (UnionString) ImplementsSettingEditResponseZonesChinaNetworkEnabledValueUnion()               {}
func (UnionString) ImplementsSettingEditResponseZonesContentConverterValueUnion()                  {}
func (UnionString) ImplementsSettingEditResponseZonesCNAMEFlatteningValueUnion()                   {}
func (UnionString) ImplementsSettingEditResponseZonesEdgeCacheTtl2ValueUnion()                     {}
func (UnionString) ImplementsSettingEditResponseZonesEmailObfuscation2ValueUnion()                 {}
func (UnionString) ImplementsSettingEditResponseZonesIPGeolocation2ValueUnion()                    {}
func (UnionString) ImplementsSettingEditResponseZonesMaxUploadValueUnion()                         {}
func (UnionString) ImplementsSettingEditResponseZonesMirage2ValueUnion()                           {}
func (UnionString) ImplementsSettingEditResponseZonesOpportunisticEncryption2ValueUnion()          {}
func (UnionString) ImplementsSettingEditResponseZonesOriginErrorPagePassThru2ValueUnion()          {}
func (UnionString) ImplementsSettingEditResponseZonesPolish2ValueUnion()                           {}
func (UnionString) ImplementsSettingEditResponseZonesPreRenderValueUnion()                         {}
func (UnionString) ImplementsSettingEditResponseZonesPrivacyPassValueUnion()                       {}
func (UnionString) ImplementsSettingEditResponseZonesRedirectsForAITrainingValueUnion()            {}
func (UnionString) ImplementsSettingEditResponseZonesReplaceInsecureJSValueUnion()                 {}
func (UnionString) ImplementsSettingEditResponseZonesResponseBuffering2ValueUnion()                {}
func (UnionString) ImplementsSettingEditResponseZonesRocketLoader2ValueUnion()                     {}
func (UnionString) ImplementsSettingEditResponseZonesSearchForAgentsValueUnion()                   {}
func (UnionString) ImplementsSettingEditResponseZonesSecurityLevel2ValueUnion()                    {}
func (UnionString) ImplementsSettingEditResponseZonesSha1SupportValueUnion()                       {}
func (UnionString) ImplementsSettingEditResponseZonesSortQueryStringForCache2ValueUnion()          {}
func (UnionString) ImplementsSettingEditResponseZonesSsl2ValueUnion()                              {}
func (UnionString) ImplementsSettingEditResponseZonesTls1_2OnlyValueUnion()                        {}
func (UnionString) ImplementsSettingEditResponseZonesTransformationsValueUnion()                   {}
func (UnionString) ImplementsSettingEditResponseZonesTransformationsAllowedOriginsValueUnion()     {}
func (UnionString) ImplementsSettingEditResponseZonesTrueClientIPHeader2ValueUnion()               {}
func (UnionString) ImplementsSettingEditResponseZonesWaf2ValueUnion()                              {}
func (UnionString) ImplementsSettingEditResponseZonesWebmcpEnabledValueUnion()                     {}
func (UnionString) ImplementsSettingEditResponseZonesWebmcpMcpURLValueUnion()                      {}
func (UnionString) ImplementsSettingEditResponseZonesWebmcpPacksValueUnion()                       {}
func (UnionString) ImplementsSettingEditResponseZonesAVIFValueUnion()                              {}
func (UnionString) ImplementsSettingGetResponseZonesAlwaysUseHttps2ValueUnion()                    {}
func (UnionString) ImplementsSettingGetResponseZonesAutomaticHTTPSRewrites2ValueUnion()            {}
func (UnionString) ImplementsSettingGetResponseZonesBrowserCacheTtl2ValueUnion()                   {}
func (UnionString) ImplementsSettingGetResponseZonesBrowserCheck2ValueUnion()                      {}
func (UnionString) ImplementsSettingGetResponseZonesCacheLevel2ValueUnion()                        {}
func (UnionString) ImplementsSettingGetResponseZonesChinaNetworkEnabledValueUnion()                {}
func (UnionString) ImplementsSettingGetResponseZonesContentConverterValueUnion()                   {}
func (UnionString) ImplementsSettingGetResponseZonesCNAMEFlatteningValueUnion()                    {}
func (UnionString) ImplementsSettingGetResponseZonesEdgeCacheTtl2ValueUnion()                      {}
func (UnionString) ImplementsSettingGetResponseZonesEmailObfuscation2ValueUnion()                  {}
func (UnionString) ImplementsSettingGetResponseZonesIPGeolocation2ValueUnion()                     {}
func (UnionString) ImplementsSettingGetResponseZonesMaxUploadValueUnion()                          {}
func (UnionString) ImplementsSettingGetResponseZonesMirage2ValueUnion()                            {}
func (UnionString) ImplementsSettingGetResponseZonesOpportunisticEncryption2ValueUnion()           {}
func (UnionString) ImplementsSettingGetResponseZonesOriginErrorPagePassThru2ValueUnion()           {}
func (UnionString) ImplementsSettingGetResponseZonesPolish2ValueUnion()                            {}
func (UnionString) ImplementsSettingGetResponseZonesPreRenderValueUnion()                          {}
func (UnionString) ImplementsSettingGetResponseZonesPrivacyPassValueUnion()                        {}
func (UnionString) ImplementsSettingGetResponseZonesRedirectsForAITrainingValueUnion()             {}
func (UnionString) ImplementsSettingGetResponseZonesReplaceInsecureJSValueUnion()                  {}
func (UnionString) ImplementsSettingGetResponseZonesResponseBuffering2ValueUnion()                 {}
func (UnionString) ImplementsSettingGetResponseZonesRocketLoader2ValueUnion()                      {}
func (UnionString) ImplementsSettingGetResponseZonesSearchForAgentsValueUnion()                    {}
func (UnionString) ImplementsSettingGetResponseZonesSecurityLevel2ValueUnion()                     {}
func (UnionString) ImplementsSettingGetResponseZonesSha1SupportValueUnion()                        {}
func (UnionString) ImplementsSettingGetResponseZonesSortQueryStringForCache2ValueUnion()           {}
func (UnionString) ImplementsSettingGetResponseZonesSsl2ValueUnion()                               {}
func (UnionString) ImplementsSettingGetResponseZonesTls1_2OnlyValueUnion()                         {}
func (UnionString) ImplementsSettingGetResponseZonesTransformationsValueUnion()                    {}
func (UnionString) ImplementsSettingGetResponseZonesTransformationsAllowedOriginsValueUnion()      {}
func (UnionString) ImplementsSettingGetResponseZonesTrueClientIPHeader2ValueUnion()                {}
func (UnionString) ImplementsSettingGetResponseZonesWaf2ValueUnion()                               {}
func (UnionString) ImplementsSettingGetResponseZonesWebmcpEnabledValueUnion()                      {}
func (UnionString) ImplementsSettingGetResponseZonesWebmcpMcpURLValueUnion()                       {}
func (UnionString) ImplementsSettingGetResponseZonesWebmcpPacksValueUnion()                        {}
func (UnionString) ImplementsSettingGetResponseZonesAVIFValueUnion()                               {}
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
func (UnionString) ImplementsObservabilityAgentTraceEventsResponseSourceUnion()   {}
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
func (UnionString) ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion()         {}
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

func (UnionBool) ImplementsEntitlementListResponseAllocationValueUnion() {}
func (UnionBool) ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion() {
}
func (UnionBool) ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion() {}
func (UnionBool) ImplementsAdvancedDDoSValueUnion()                                                 {}
func (UnionBool) ImplementsAlwaysOnlineValueUnion()                                                 {}
func (UnionBool) ImplementsBrotliValueUnion()                                                       {}
func (UnionBool) ImplementsChallengeTTLValueUnion()                                                 {}
func (UnionBool) ImplementsCiphersValueUnion()                                                      {}
func (UnionBool) ImplementsDevelopmentModeValueUnion()                                              {}
func (UnionBool) ImplementsEarlyHintsValueUnion()                                                   {}
func (UnionBool) ImplementsH2PrioritizationValueUnion()                                             {}
func (UnionBool) ImplementsHotlinkProtectionValueUnion()                                            {}
func (UnionBool) ImplementsHTTP2ValueUnion()                                                        {}
func (UnionBool) ImplementsHTTP3ValueUnion()                                                        {}
func (UnionBool) ImplementsImageResizingValueUnion()                                                {}
func (UnionBool) ImplementsIPV6ValueUnion()                                                         {}
func (UnionBool) ImplementsMinTLSVersionValueUnion()                                                {}
func (UnionBool) ImplementsOpportunisticOnionValueUnion()                                           {}
func (UnionBool) ImplementsOrangeToOrangeValueUnion()                                               {}
func (UnionBool) ImplementsPrefetchPreloadValueUnion()                                              {}
func (UnionBool) ImplementsProxyReadTimeoutValueUnion()                                             {}
func (UnionBool) ImplementsPseudoIPV4ValueUnion()                                                   {}
func (UnionBool) ImplementsServerSideExcludesValueUnion()                                           {}
func (UnionBool) ImplementsTls1_3ValueUnion()                                                       {}
func (UnionBool) ImplementsTLSClientAuthValueUnion()                                                {}
func (UnionBool) ImplementsWebPValueUnion()                                                         {}
func (UnionBool) ImplementsWebsocketValueUnion()                                                    {}
func (UnionBool) ImplementsZeroRTTValueUnion()                                                      {}
func (UnionBool) ImplementsSettingEditResponseZonesAlwaysUseHttps2ValueUnion()                      {}
func (UnionBool) ImplementsSettingEditResponseZonesAutomaticHTTPSRewrites2ValueUnion()              {}
func (UnionBool) ImplementsSettingEditResponseZonesBrowserCacheTtl2ValueUnion()                     {}
func (UnionBool) ImplementsSettingEditResponseZonesBrowserCheck2ValueUnion()                        {}
func (UnionBool) ImplementsSettingEditResponseZonesCacheLevel2ValueUnion()                          {}
func (UnionBool) ImplementsSettingEditResponseZonesChinaNetworkEnabledValueUnion()                  {}
func (UnionBool) ImplementsSettingEditResponseZonesContentConverterValueUnion()                     {}
func (UnionBool) ImplementsSettingEditResponseZonesCNAMEFlatteningValueUnion()                      {}
func (UnionBool) ImplementsSettingEditResponseZonesEdgeCacheTtl2ValueUnion()                        {}
func (UnionBool) ImplementsSettingEditResponseZonesEmailObfuscation2ValueUnion()                    {}
func (UnionBool) ImplementsSettingEditResponseZonesIPGeolocation2ValueUnion()                       {}
func (UnionBool) ImplementsSettingEditResponseZonesMaxUploadValueUnion()                            {}
func (UnionBool) ImplementsSettingEditResponseZonesMirage2ValueUnion()                              {}
func (UnionBool) ImplementsSettingEditResponseZonesOpportunisticEncryption2ValueUnion()             {}
func (UnionBool) ImplementsSettingEditResponseZonesOriginErrorPagePassThru2ValueUnion()             {}
func (UnionBool) ImplementsSettingEditResponseZonesPolish2ValueUnion()                              {}
func (UnionBool) ImplementsSettingEditResponseZonesPreRenderValueUnion()                            {}
func (UnionBool) ImplementsSettingEditResponseZonesPrivacyPassValueUnion()                          {}
func (UnionBool) ImplementsSettingEditResponseZonesRedirectsForAITrainingValueUnion()               {}
func (UnionBool) ImplementsSettingEditResponseZonesReplaceInsecureJSValueUnion()                    {}
func (UnionBool) ImplementsSettingEditResponseZonesResponseBuffering2ValueUnion()                   {}
func (UnionBool) ImplementsSettingEditResponseZonesRocketLoader2ValueUnion()                        {}
func (UnionBool) ImplementsSettingEditResponseZonesSearchForAgentsValueUnion()                      {}
func (UnionBool) ImplementsSettingEditResponseZonesSecurityLevel2ValueUnion()                       {}
func (UnionBool) ImplementsSettingEditResponseZonesSha1SupportValueUnion()                          {}
func (UnionBool) ImplementsSettingEditResponseZonesSortQueryStringForCache2ValueUnion()             {}
func (UnionBool) ImplementsSettingEditResponseZonesSsl2ValueUnion()                                 {}
func (UnionBool) ImplementsSettingEditResponseZonesTls1_2OnlyValueUnion()                           {}
func (UnionBool) ImplementsSettingEditResponseZonesTransformationsValueUnion()                      {}
func (UnionBool) ImplementsSettingEditResponseZonesTransformationsAllowedOriginsValueUnion()        {}
func (UnionBool) ImplementsSettingEditResponseZonesTrueClientIPHeader2ValueUnion()                  {}
func (UnionBool) ImplementsSettingEditResponseZonesWaf2ValueUnion()                                 {}
func (UnionBool) ImplementsSettingEditResponseZonesWebmcpEnabledValueUnion()                        {}
func (UnionBool) ImplementsSettingEditResponseZonesWebmcpMcpURLValueUnion()                         {}
func (UnionBool) ImplementsSettingEditResponseZonesWebmcpPacksValueUnion()                          {}
func (UnionBool) ImplementsSettingEditResponseZonesAVIFValueUnion()                                 {}
func (UnionBool) ImplementsSettingGetResponseZonesAlwaysUseHttps2ValueUnion()                       {}
func (UnionBool) ImplementsSettingGetResponseZonesAutomaticHTTPSRewrites2ValueUnion()               {}
func (UnionBool) ImplementsSettingGetResponseZonesBrowserCacheTtl2ValueUnion()                      {}
func (UnionBool) ImplementsSettingGetResponseZonesBrowserCheck2ValueUnion()                         {}
func (UnionBool) ImplementsSettingGetResponseZonesCacheLevel2ValueUnion()                           {}
func (UnionBool) ImplementsSettingGetResponseZonesChinaNetworkEnabledValueUnion()                   {}
func (UnionBool) ImplementsSettingGetResponseZonesContentConverterValueUnion()                      {}
func (UnionBool) ImplementsSettingGetResponseZonesCNAMEFlatteningValueUnion()                       {}
func (UnionBool) ImplementsSettingGetResponseZonesEdgeCacheTtl2ValueUnion()                         {}
func (UnionBool) ImplementsSettingGetResponseZonesEmailObfuscation2ValueUnion()                     {}
func (UnionBool) ImplementsSettingGetResponseZonesIPGeolocation2ValueUnion()                        {}
func (UnionBool) ImplementsSettingGetResponseZonesMaxUploadValueUnion()                             {}
func (UnionBool) ImplementsSettingGetResponseZonesMirage2ValueUnion()                               {}
func (UnionBool) ImplementsSettingGetResponseZonesOpportunisticEncryption2ValueUnion()              {}
func (UnionBool) ImplementsSettingGetResponseZonesOriginErrorPagePassThru2ValueUnion()              {}
func (UnionBool) ImplementsSettingGetResponseZonesPolish2ValueUnion()                               {}
func (UnionBool) ImplementsSettingGetResponseZonesPreRenderValueUnion()                             {}
func (UnionBool) ImplementsSettingGetResponseZonesPrivacyPassValueUnion()                           {}
func (UnionBool) ImplementsSettingGetResponseZonesRedirectsForAITrainingValueUnion()                {}
func (UnionBool) ImplementsSettingGetResponseZonesReplaceInsecureJSValueUnion()                     {}
func (UnionBool) ImplementsSettingGetResponseZonesResponseBuffering2ValueUnion()                    {}
func (UnionBool) ImplementsSettingGetResponseZonesRocketLoader2ValueUnion()                         {}
func (UnionBool) ImplementsSettingGetResponseZonesSearchForAgentsValueUnion()                       {}
func (UnionBool) ImplementsSettingGetResponseZonesSecurityLevel2ValueUnion()                        {}
func (UnionBool) ImplementsSettingGetResponseZonesSha1SupportValueUnion()                           {}
func (UnionBool) ImplementsSettingGetResponseZonesSortQueryStringForCache2ValueUnion()              {}
func (UnionBool) ImplementsSettingGetResponseZonesSsl2ValueUnion()                                  {}
func (UnionBool) ImplementsSettingGetResponseZonesTls1_2OnlyValueUnion()                            {}
func (UnionBool) ImplementsSettingGetResponseZonesTransformationsValueUnion()                       {}
func (UnionBool) ImplementsSettingGetResponseZonesTransformationsAllowedOriginsValueUnion()         {}
func (UnionBool) ImplementsSettingGetResponseZonesTrueClientIPHeader2ValueUnion()                   {}
func (UnionBool) ImplementsSettingGetResponseZonesWaf2ValueUnion()                                  {}
func (UnionBool) ImplementsSettingGetResponseZonesWebmcpEnabledValueUnion()                         {}
func (UnionBool) ImplementsSettingGetResponseZonesWebmcpMcpURLValueUnion()                          {}
func (UnionBool) ImplementsSettingGetResponseZonesWebmcpPacksValueUnion()                           {}
func (UnionBool) ImplementsSettingGetResponseZonesAVIFValueUnion()                                  {}
func (UnionBool) ImplementsAnalyticsQuerySummaryParamsFiltersValueUnion()                           {}
func (UnionBool) ImplementsAnalyticsQueryTimeseriesParamsFiltersValueUnion()                        {}
func (UnionBool) ImplementsAnalyticsQueryTopNParamsFiltersValueUnion()                              {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion()    {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion()        {}
func (UnionBool) ImplementsAnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion()     {}
func (UnionBool) ImplementsVersionAssetsConfigRunWorkerFirstUnionParam()                            {}
func (UnionBool) ImplementsVersionAssetsConfigRunWorkerFirstUnion()                                 {}
func (UnionBool) ImplementsScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()              {}
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
func (UnionBool) ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion()         {}
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

func (UnionFloat) ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion() {
}
func (UnionFloat) ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion() {
}
func (UnionFloat) ImplementsAdvancedDDoSValueUnion()                                              {}
func (UnionFloat) ImplementsAlwaysOnlineValueUnion()                                              {}
func (UnionFloat) ImplementsBrotliValueUnion()                                                    {}
func (UnionFloat) ImplementsChallengeTTLValueUnion()                                              {}
func (UnionFloat) ImplementsCiphersValueUnion()                                                   {}
func (UnionFloat) ImplementsDevelopmentModeValueUnion()                                           {}
func (UnionFloat) ImplementsEarlyHintsValueUnion()                                                {}
func (UnionFloat) ImplementsH2PrioritizationValueUnion()                                          {}
func (UnionFloat) ImplementsHotlinkProtectionValueUnion()                                         {}
func (UnionFloat) ImplementsHTTP2ValueUnion()                                                     {}
func (UnionFloat) ImplementsHTTP3ValueUnion()                                                     {}
func (UnionFloat) ImplementsImageResizingValueUnion()                                             {}
func (UnionFloat) ImplementsIPV6ValueUnion()                                                      {}
func (UnionFloat) ImplementsMinTLSVersionValueUnion()                                             {}
func (UnionFloat) ImplementsOpportunisticOnionValueUnion()                                        {}
func (UnionFloat) ImplementsOrangeToOrangeValueUnion()                                            {}
func (UnionFloat) ImplementsPrefetchPreloadValueUnion()                                           {}
func (UnionFloat) ImplementsProxyReadTimeoutValueUnion()                                          {}
func (UnionFloat) ImplementsPseudoIPV4ValueUnion()                                                {}
func (UnionFloat) ImplementsServerSideExcludesValueUnion()                                        {}
func (UnionFloat) ImplementsTls1_3ValueUnion()                                                    {}
func (UnionFloat) ImplementsTLSClientAuthValueUnion()                                             {}
func (UnionFloat) ImplementsWebPValueUnion()                                                      {}
func (UnionFloat) ImplementsWebsocketValueUnion()                                                 {}
func (UnionFloat) ImplementsZeroRTTValueUnion()                                                   {}
func (UnionFloat) ImplementsSettingEditResponseZonesAlwaysUseHttps2ValueUnion()                   {}
func (UnionFloat) ImplementsSettingEditResponseZonesAutomaticHTTPSRewrites2ValueUnion()           {}
func (UnionFloat) ImplementsSettingEditResponseZonesBrowserCacheTtl2ValueUnion()                  {}
func (UnionFloat) ImplementsSettingEditResponseZonesBrowserCheck2ValueUnion()                     {}
func (UnionFloat) ImplementsSettingEditResponseZonesCacheLevel2ValueUnion()                       {}
func (UnionFloat) ImplementsSettingEditResponseZonesChinaNetworkEnabledValueUnion()               {}
func (UnionFloat) ImplementsSettingEditResponseZonesContentConverterValueUnion()                  {}
func (UnionFloat) ImplementsSettingEditResponseZonesCNAMEFlatteningValueUnion()                   {}
func (UnionFloat) ImplementsSettingEditResponseZonesEdgeCacheTtl2ValueUnion()                     {}
func (UnionFloat) ImplementsSettingEditResponseZonesEmailObfuscation2ValueUnion()                 {}
func (UnionFloat) ImplementsSettingEditResponseZonesIPGeolocation2ValueUnion()                    {}
func (UnionFloat) ImplementsSettingEditResponseZonesMaxUploadValueUnion()                         {}
func (UnionFloat) ImplementsSettingEditResponseZonesMirage2ValueUnion()                           {}
func (UnionFloat) ImplementsSettingEditResponseZonesOpportunisticEncryption2ValueUnion()          {}
func (UnionFloat) ImplementsSettingEditResponseZonesOriginErrorPagePassThru2ValueUnion()          {}
func (UnionFloat) ImplementsSettingEditResponseZonesPolish2ValueUnion()                           {}
func (UnionFloat) ImplementsSettingEditResponseZonesPreRenderValueUnion()                         {}
func (UnionFloat) ImplementsSettingEditResponseZonesPrivacyPassValueUnion()                       {}
func (UnionFloat) ImplementsSettingEditResponseZonesRedirectsForAITrainingValueUnion()            {}
func (UnionFloat) ImplementsSettingEditResponseZonesReplaceInsecureJSValueUnion()                 {}
func (UnionFloat) ImplementsSettingEditResponseZonesResponseBuffering2ValueUnion()                {}
func (UnionFloat) ImplementsSettingEditResponseZonesRocketLoader2ValueUnion()                     {}
func (UnionFloat) ImplementsSettingEditResponseZonesSearchForAgentsValueUnion()                   {}
func (UnionFloat) ImplementsSettingEditResponseZonesSecurityLevel2ValueUnion()                    {}
func (UnionFloat) ImplementsSettingEditResponseZonesSha1SupportValueUnion()                       {}
func (UnionFloat) ImplementsSettingEditResponseZonesSortQueryStringForCache2ValueUnion()          {}
func (UnionFloat) ImplementsSettingEditResponseZonesSsl2ValueUnion()                              {}
func (UnionFloat) ImplementsSettingEditResponseZonesTls1_2OnlyValueUnion()                        {}
func (UnionFloat) ImplementsSettingEditResponseZonesTransformationsValueUnion()                   {}
func (UnionFloat) ImplementsSettingEditResponseZonesTransformationsAllowedOriginsValueUnion()     {}
func (UnionFloat) ImplementsSettingEditResponseZonesTrueClientIPHeader2ValueUnion()               {}
func (UnionFloat) ImplementsSettingEditResponseZonesWaf2ValueUnion()                              {}
func (UnionFloat) ImplementsSettingEditResponseZonesWebmcpEnabledValueUnion()                     {}
func (UnionFloat) ImplementsSettingEditResponseZonesWebmcpMcpURLValueUnion()                      {}
func (UnionFloat) ImplementsSettingEditResponseZonesWebmcpPacksValueUnion()                       {}
func (UnionFloat) ImplementsSettingEditResponseZonesAVIFValueUnion()                              {}
func (UnionFloat) ImplementsSettingGetResponseZonesAlwaysUseHttps2ValueUnion()                    {}
func (UnionFloat) ImplementsSettingGetResponseZonesAutomaticHTTPSRewrites2ValueUnion()            {}
func (UnionFloat) ImplementsSettingGetResponseZonesBrowserCacheTtl2ValueUnion()                   {}
func (UnionFloat) ImplementsSettingGetResponseZonesBrowserCheck2ValueUnion()                      {}
func (UnionFloat) ImplementsSettingGetResponseZonesCacheLevel2ValueUnion()                        {}
func (UnionFloat) ImplementsSettingGetResponseZonesChinaNetworkEnabledValueUnion()                {}
func (UnionFloat) ImplementsSettingGetResponseZonesContentConverterValueUnion()                   {}
func (UnionFloat) ImplementsSettingGetResponseZonesCNAMEFlatteningValueUnion()                    {}
func (UnionFloat) ImplementsSettingGetResponseZonesEdgeCacheTtl2ValueUnion()                      {}
func (UnionFloat) ImplementsSettingGetResponseZonesEmailObfuscation2ValueUnion()                  {}
func (UnionFloat) ImplementsSettingGetResponseZonesIPGeolocation2ValueUnion()                     {}
func (UnionFloat) ImplementsSettingGetResponseZonesMaxUploadValueUnion()                          {}
func (UnionFloat) ImplementsSettingGetResponseZonesMirage2ValueUnion()                            {}
func (UnionFloat) ImplementsSettingGetResponseZonesOpportunisticEncryption2ValueUnion()           {}
func (UnionFloat) ImplementsSettingGetResponseZonesOriginErrorPagePassThru2ValueUnion()           {}
func (UnionFloat) ImplementsSettingGetResponseZonesPolish2ValueUnion()                            {}
func (UnionFloat) ImplementsSettingGetResponseZonesPreRenderValueUnion()                          {}
func (UnionFloat) ImplementsSettingGetResponseZonesPrivacyPassValueUnion()                        {}
func (UnionFloat) ImplementsSettingGetResponseZonesRedirectsForAITrainingValueUnion()             {}
func (UnionFloat) ImplementsSettingGetResponseZonesReplaceInsecureJSValueUnion()                  {}
func (UnionFloat) ImplementsSettingGetResponseZonesResponseBuffering2ValueUnion()                 {}
func (UnionFloat) ImplementsSettingGetResponseZonesRocketLoader2ValueUnion()                      {}
func (UnionFloat) ImplementsSettingGetResponseZonesSearchForAgentsValueUnion()                    {}
func (UnionFloat) ImplementsSettingGetResponseZonesSecurityLevel2ValueUnion()                     {}
func (UnionFloat) ImplementsSettingGetResponseZonesSha1SupportValueUnion()                        {}
func (UnionFloat) ImplementsSettingGetResponseZonesSortQueryStringForCache2ValueUnion()           {}
func (UnionFloat) ImplementsSettingGetResponseZonesSsl2ValueUnion()                               {}
func (UnionFloat) ImplementsSettingGetResponseZonesTls1_2OnlyValueUnion()                         {}
func (UnionFloat) ImplementsSettingGetResponseZonesTransformationsValueUnion()                    {}
func (UnionFloat) ImplementsSettingGetResponseZonesTransformationsAllowedOriginsValueUnion()      {}
func (UnionFloat) ImplementsSettingGetResponseZonesTrueClientIPHeader2ValueUnion()                {}
func (UnionFloat) ImplementsSettingGetResponseZonesWaf2ValueUnion()                               {}
func (UnionFloat) ImplementsSettingGetResponseZonesWebmcpEnabledValueUnion()                      {}
func (UnionFloat) ImplementsSettingGetResponseZonesWebmcpMcpURLValueUnion()                       {}
func (UnionFloat) ImplementsSettingGetResponseZonesWebmcpPacksValueUnion()                        {}
func (UnionFloat) ImplementsSettingGetResponseZonesAVIFValueUnion()                               {}
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
func (UnionFloat) ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion()         {}
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

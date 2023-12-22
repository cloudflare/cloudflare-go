// File generated from our OpenAPI spec by Stainless.

package shared

type UnionString string

func (UnionString) ImplementsAPIResponseSingleResult()                                          {}
func (UnionString) ImplementsDNSFirewallResponseCollectionResultDNSFirewallIP()                 {}
func (UnionString) ImplementsDNSFirewallResponseCollectionResultOriginIP()                      {}
func (UnionString) ImplementsDNSFirewallSingleResponseResultDNSFirewallIP()                     {}
func (UnionString) ImplementsDNSFirewallSingleResponseResultOriginIP()                          {}
func (UnionString) ImplementsAccountDNSFirewallUpdateParamsDNSFirewallIP()                      {}
func (UnionString) ImplementsAccountDNSFirewallUpdateParamsOriginIP()                           {}
func (UnionString) ImplementsAccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsOriginIP() {}
func (UnionString) ImplementsImageResponseCollectionResultVariant()                             {}
func (UnionString) ImplementsImageResponseBlob()                                                {}
func (UnionString) ImplementsSchemasResponseResultIP()                                          {}
func (UnionString) ImplementsPreviewResultResponseResult()                                      {}
func (UnionString) ImplementsDeploymentsListResponseResult()                                    {}
func (UnionString) ImplementsDeploymentsSingleResponseResult()                                  {}
func (UnionString) ImplementsAPIResponseCommonResult()                                          {}
func (UnionString) ImplementsKeysSingleResponseResult()                                         {}
func (UnionString) ImplementsZonePurgeCachZonePurgeParamsFilesFile()                            {}
func (UnionString) ImplementsColoResponseQuerySince()                                           {}
func (UnionString) ImplementsColoResponseQueryUntil()                                           {}
func (UnionString) ImplementsColoResponseResultTimeseriesSince()                                {}
func (UnionString) ImplementsColoResponseResultTimeseriesUntil()                                {}
func (UnionString) ImplementsColoResponseResultTotalsSince()                                    {}
func (UnionString) ImplementsColoResponseResultTotalsUntil()                                    {}
func (UnionString) ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince() {
}
func (UnionString) ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil() {
}
func (UnionString) ImplementsDashboardResponseQuerySince()                                          {}
func (UnionString) ImplementsDashboardResponseQueryUntil()                                          {}
func (UnionString) ImplementsDashboardResponseResultTimeseriesSince()                               {}
func (UnionString) ImplementsDashboardResponseResultTimeseriesUntil()                               {}
func (UnionString) ImplementsDashboardResponseResultTotalsSince()                                   {}
func (UnionString) ImplementsDashboardResponseResultTotalsUntil()                                   {}
func (UnionString) ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince() {}
func (UnionString) ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil() {}
func (UnionString) ImplementsLog()                                                                  {}
func (UnionString) ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsEnd()                      {}
func (UnionString) ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsStart()                    {}
func (UnionString) ImplementsAIModelTextEmbeddingBgeSmallEnV1_5RunParamsText()                      {}
func (UnionString) ImplementsAIModelTextEmbeddingBgeBaseEnV1_5RunParamsText()                       {}
func (UnionString) ImplementsAIModelTextEmbeddingBgeLargeEnV1_5RunParamsText()                      {}
func (UnionString) ImplementsAIModelTextGenerationLlama2_7bChatInt8RunResponse()                    {}
func (UnionString) ImplementsAIModelTextGenerationLlama2_7bChatFp16RunResponse()                    {}
func (UnionString) ImplementsAIModelTextGenerationMistral7bInstructV0_1RunResponse()                {}

type UnionInt int64

func (UnionInt) ImplementsColoResponseQuerySince()            {}
func (UnionInt) ImplementsColoResponseQueryUntil()            {}
func (UnionInt) ImplementsColoResponseResultTimeseriesSince() {}
func (UnionInt) ImplementsColoResponseResultTimeseriesUntil() {}
func (UnionInt) ImplementsColoResponseResultTotalsSince()     {}
func (UnionInt) ImplementsColoResponseResultTotalsUntil()     {}
func (UnionInt) ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince() {
}
func (UnionInt) ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil() {
}
func (UnionInt) ImplementsDashboardResponseQuerySince()                                          {}
func (UnionInt) ImplementsDashboardResponseQueryUntil()                                          {}
func (UnionInt) ImplementsDashboardResponseResultTimeseriesSince()                               {}
func (UnionInt) ImplementsDashboardResponseResultTimeseriesUntil()                               {}
func (UnionInt) ImplementsDashboardResponseResultTotalsSince()                                   {}
func (UnionInt) ImplementsDashboardResponseResultTotalsUntil()                                   {}
func (UnionInt) ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince() {}
func (UnionInt) ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil() {}
func (UnionInt) ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsEnd()                      {}
func (UnionInt) ImplementsZoneLogReceivedReceivedGetLogsReceivedParamsStart()                    {}

type UnionFloat float64

func (UnionFloat) ImplementsDNSResponseCollectionResultARecordTtl()                            {}
func (UnionFloat) ImplementsDNSResponseCollectionResultAaaaRecordTtl()                         {}
func (UnionFloat) ImplementsDNSResponseCollectionResultCaaRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseCollectionResultCertRecordTtl()                         {}
func (UnionFloat) ImplementsDNSResponseCollectionResultCnameRecordTtl()                        {}
func (UnionFloat) ImplementsDNSResponseCollectionResultDnskeyRecordTtl()                       {}
func (UnionFloat) ImplementsDNSResponseCollectionResultDsRecordTtl()                           {}
func (UnionFloat) ImplementsDNSResponseCollectionResultHTTPsRecordTtl()                        {}
func (UnionFloat) ImplementsDNSResponseCollectionResultLocRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseCollectionResultMxRecordTtl()                           {}
func (UnionFloat) ImplementsDNSResponseCollectionResultNaptrRecordTtl()                        {}
func (UnionFloat) ImplementsDNSResponseCollectionResultNsRecordTtl()                           {}
func (UnionFloat) ImplementsDNSResponseCollectionResultPtrRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseCollectionResultSmimeaRecordTtl()                       {}
func (UnionFloat) ImplementsDNSResponseCollectionResultSrvRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseCollectionResultSshfpRecordTtl()                        {}
func (UnionFloat) ImplementsDNSResponseCollectionResultSvcbRecordTtl()                         {}
func (UnionFloat) ImplementsDNSResponseCollectionResultTlsaRecordTtl()                         {}
func (UnionFloat) ImplementsDNSResponseCollectionResultTxtRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseCollectionResultUriRecordTtl()                          {}
func (UnionFloat) ImplementsDNSResponseSingleResultARecordTtl()                                {}
func (UnionFloat) ImplementsDNSResponseSingleResultAaaaRecordTtl()                             {}
func (UnionFloat) ImplementsDNSResponseSingleResultCaaRecordTtl()                              {}
func (UnionFloat) ImplementsDNSResponseSingleResultCertRecordTtl()                             {}
func (UnionFloat) ImplementsDNSResponseSingleResultCnameRecordTtl()                            {}
func (UnionFloat) ImplementsDNSResponseSingleResultDnskeyRecordTtl()                           {}
func (UnionFloat) ImplementsDNSResponseSingleResultDsRecordTtl()                               {}
func (UnionFloat) ImplementsDNSResponseSingleResultHTTPsRecordTtl()                            {}
func (UnionFloat) ImplementsDNSResponseSingleResultLocRecordTtl()                              {}
func (UnionFloat) ImplementsDNSResponseSingleResultMxRecordTtl()                               {}
func (UnionFloat) ImplementsDNSResponseSingleResultNaptrRecordTtl()                            {}
func (UnionFloat) ImplementsDNSResponseSingleResultNsRecordTtl()                               {}
func (UnionFloat) ImplementsDNSResponseSingleResultPtrRecordTtl()                              {}
func (UnionFloat) ImplementsDNSResponseSingleResultSmimeaRecordTtl()                           {}
func (UnionFloat) ImplementsDNSResponseSingleResultSrvRecordTtl()                              {}
func (UnionFloat) ImplementsDNSResponseSingleResultSshfpRecordTtl()                            {}
func (UnionFloat) ImplementsDNSResponseSingleResultSvcbRecordTtl()                             {}
func (UnionFloat) ImplementsDNSResponseSingleResultTlsaRecordTtl()                             {}
func (UnionFloat) ImplementsDNSResponseSingleResultTxtRecordTtl()                              {}
func (UnionFloat) ImplementsDNSResponseSingleResultUriRecordTtl()                              {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsARecordTtl()                              {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsAaaaRecordTtl()                           {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsCaaRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsCertRecordTtl()                           {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsCnameRecordTtl()                          {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsDnskeyRecordTtl()                         {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsDsRecordTtl()                             {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsHTTPsRecordTtl()                          {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsLocRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsMxRecordTtl()                             {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsNaptrRecordTtl()                          {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsNsRecordTtl()                             {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsPtrRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsSmimeaRecordTtl()                         {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsSrvRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsSshfpRecordTtl()                          {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsSvcbRecordTtl()                           {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsTlsaRecordTtl()                           {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsTxtRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordUpdateParamsUriRecordTtl()                            {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtl()      {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtl()   {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtl()    {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtl()   {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtl()  {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtl() {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtl()     {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtl()  {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtl()    {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtl()     {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtl()  {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtl()     {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtl()    {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtl() {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtl()    {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtl()  {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtl()   {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtl()   {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtl()    {}
func (UnionFloat) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtl()    {}
func (UnionFloat) ImplementsDNSSettingsResponseCollectionResultTtl()                           {}

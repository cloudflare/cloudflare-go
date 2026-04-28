// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_search"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/r2"
)

func TestInstanceNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.New(context.TODO(), ai_search.InstanceNewParams{
		AccountID:      cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		ID:             cloudflare.F("my-ai-search"),
		AIGatewayID:    cloudflare.F("ai_gateway_id"),
		AISearchModel:  cloudflare.F(ai_search.InstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
		Cache:          cloudflare.F(true),
		CacheThreshold: cloudflare.F(ai_search.InstanceNewParamsCacheThresholdSuperStrictMatch),
		Chunk:          cloudflare.F(true),
		ChunkOverlap:   cloudflare.F(int64(0)),
		ChunkSize:      cloudflare.F(int64(64)),
		CustomMetadata: cloudflare.F([]ai_search.InstanceNewParamsCustomMetadata{{
			DataType:  cloudflare.F(ai_search.InstanceNewParamsCustomMetadataDataTypeText),
			FieldName: cloudflare.F("x"),
		}}),
		EmbeddingModel:      cloudflare.F(ai_search.InstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
		FusionMethod:        cloudflare.F(ai_search.InstanceNewParamsFusionMethodMax),
		HybridSearchEnabled: cloudflare.F(true),
		IndexMethod: cloudflare.F(ai_search.InstanceNewParamsIndexMethod{
			Keyword: cloudflare.F(true),
			Vector:  cloudflare.F(true),
		}),
		IndexingOptions: cloudflare.F(ai_search.InstanceNewParamsIndexingOptions{
			KeywordTokenizer: cloudflare.F(ai_search.InstanceNewParamsIndexingOptionsKeywordTokenizerPorter),
		}),
		MaxNumResults: cloudflare.F(int64(1)),
		Metadata: cloudflare.F(ai_search.InstanceNewParamsMetadata{
			CreatedFromAISearchWizard: cloudflare.F(true),
			SearchForAgents: cloudflare.F(ai_search.InstanceNewParamsMetadataSearchForAgents{
				Hostname: cloudflare.F("hostname"),
				ZoneID:   cloudflare.F("zone_id"),
				ZoneName: cloudflare.F("zone_name"),
			}),
			WorkerDomain: cloudflare.F("worker_domain"),
		}),
		PublicEndpointParams: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParams{
			AuthorizedHosts: cloudflare.F([]string{"string"}),
			ChatCompletionsEndpoint: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint{
				Disabled: cloudflare.F(true),
			}),
			Enabled: cloudflare.F(true),
			Mcp: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsMcp{
				Description: cloudflare.F("description"),
				Disabled:    cloudflare.F(true),
			}),
			RateLimit: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsRateLimit{
				PeriodMs:  cloudflare.F(int64(60000)),
				Requests:  cloudflare.F(int64(1)),
				Technique: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed),
			}),
			SearchEndpoint: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsSearchEndpoint{
				Disabled: cloudflare.F(true),
			}),
		}),
		Reranking:      cloudflare.F(true),
		RerankingModel: cloudflare.F(ai_search.InstanceNewParamsRerankingModelCfBaaiBgeRerankerBase),
		RetrievalOptions: cloudflare.F(ai_search.InstanceNewParamsRetrievalOptions{
			BoostBy: cloudflare.F([]ai_search.InstanceNewParamsRetrievalOptionsBoostBy{{
				Field:     cloudflare.F("timestamp"),
				Direction: cloudflare.F(ai_search.InstanceNewParamsRetrievalOptionsBoostByDirectionDesc),
			}}),
			KeywordMatchMode: cloudflare.F(ai_search.InstanceNewParamsRetrievalOptionsKeywordMatchModeAnd),
		}),
		RewriteModel:   cloudflare.F(ai_search.InstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
		RewriteQuery:   cloudflare.F(true),
		ScoreThreshold: cloudflare.F(0.000000),
		Source:         cloudflare.F("source"),
		SourceParams: cloudflare.F(ai_search.InstanceNewParamsSourceParams{
			ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
			IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
			Prefix:         cloudflare.F("prefix"),
			R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
			WebCrawler: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawler{
				CrawlOptions: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerCrawlOptions{
					Depth:                cloudflare.F(1.000000),
					IncludeExternalLinks: cloudflare.F(true),
					IncludeSubdomains:    cloudflare.F(true),
					MaxAge:               cloudflare.F(0.000000),
					Source:               cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceAll),
				}),
				ParseOptions: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerParseOptions{
					ContentSelector: cloudflare.F([]ai_search.InstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector{{
						Path:     cloudflare.F("**/blog/**"),
						Selector: cloudflare.F("article .post-body"),
					}}),
					IncludeHeaders: cloudflare.F(map[string]string{
						"foo": "string",
					}),
					IncludeImages:       cloudflare.F(true),
					SpecificSitemaps:    cloudflare.F([]string{"https://example.com/sitemap.xml", "https://example.com/blog-sitemap.xml"}),
					UseBrowserRendering: cloudflare.F(true),
				}),
				ParseType: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap),
				StoreOptions: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerStoreOptions{
					StorageID:      cloudflare.F("storage_id"),
					R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
					StorageType:    cloudflare.F(r2.ProviderR2),
				}),
			}),
		}),
		SyncInterval: cloudflare.F(ai_search.InstanceNewParamsSyncInterval900),
		TokenID:      cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:         cloudflare.F(ai_search.InstanceNewParamsTypeR2),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Update(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceUpdateParams{
			AccountID:      cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AIGatewayID:    cloudflare.F("ai_gateway_id"),
			AISearchModel:  cloudflare.F(ai_search.InstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
			Cache:          cloudflare.F(true),
			CacheThreshold: cloudflare.F(ai_search.InstanceUpdateParamsCacheThresholdSuperStrictMatch),
			Chunk:          cloudflare.F(true),
			ChunkOverlap:   cloudflare.F(int64(0)),
			ChunkSize:      cloudflare.F(int64(64)),
			CustomMetadata: cloudflare.F([]ai_search.InstanceUpdateParamsCustomMetadata{{
				DataType:  cloudflare.F(ai_search.InstanceUpdateParamsCustomMetadataDataTypeText),
				FieldName: cloudflare.F("x"),
			}}),
			EmbeddingModel: cloudflare.F(ai_search.InstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
			FusionMethod:   cloudflare.F(ai_search.InstanceUpdateParamsFusionMethodMax),
			IndexMethod: cloudflare.F(ai_search.InstanceUpdateParamsIndexMethod{
				Keyword: cloudflare.F(true),
				Vector:  cloudflare.F(true),
			}),
			IndexingOptions: cloudflare.F(ai_search.InstanceUpdateParamsIndexingOptions{
				KeywordTokenizer: cloudflare.F(ai_search.InstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter),
			}),
			MaxNumResults: cloudflare.F(int64(1)),
			Metadata: cloudflare.F(ai_search.InstanceUpdateParamsMetadata{
				CreatedFromAISearchWizard: cloudflare.F(true),
				SearchForAgents: cloudflare.F(ai_search.InstanceUpdateParamsMetadataSearchForAgents{
					Hostname: cloudflare.F("hostname"),
					ZoneID:   cloudflare.F("zone_id"),
					ZoneName: cloudflare.F("zone_name"),
				}),
				WorkerDomain: cloudflare.F("worker_domain"),
			}),
			Paused: cloudflare.F(true),
			PublicEndpointParams: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParams{
				AuthorizedHosts: cloudflare.F([]string{"string"}),
				ChatCompletionsEndpoint: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint{
					Disabled: cloudflare.F(true),
				}),
				Enabled: cloudflare.F(true),
				Mcp: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsMcp{
					Description: cloudflare.F("description"),
					Disabled:    cloudflare.F(true),
				}),
				RateLimit: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsRateLimit{
					PeriodMs:  cloudflare.F(int64(60000)),
					Requests:  cloudflare.F(int64(1)),
					Technique: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed),
				}),
				SearchEndpoint: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsSearchEndpoint{
					Disabled: cloudflare.F(true),
				}),
			}),
			Reranking:      cloudflare.F(true),
			RerankingModel: cloudflare.F(ai_search.InstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase),
			RetrievalOptions: cloudflare.F(ai_search.InstanceUpdateParamsRetrievalOptions{
				BoostBy: cloudflare.F([]ai_search.InstanceUpdateParamsRetrievalOptionsBoostBy{{
					Field:     cloudflare.F("timestamp"),
					Direction: cloudflare.F(ai_search.InstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc),
				}}),
				KeywordMatchMode: cloudflare.F(ai_search.InstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd),
			}),
			RewriteModel:   cloudflare.F(ai_search.InstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
			RewriteQuery:   cloudflare.F(true),
			ScoreThreshold: cloudflare.F(0.000000),
			SourceParams: cloudflare.F(ai_search.InstanceUpdateParamsSourceParams{
				ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
				IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
				Prefix:         cloudflare.F("prefix"),
				R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
				WebCrawler: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawler{
					CrawlOptions: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerCrawlOptions{
						Depth:                cloudflare.F(1.000000),
						IncludeExternalLinks: cloudflare.F(true),
						IncludeSubdomains:    cloudflare.F(true),
						MaxAge:               cloudflare.F(0.000000),
						Source:               cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceAll),
					}),
					ParseOptions: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerParseOptions{
						ContentSelector: cloudflare.F([]ai_search.InstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector{{
							Path:     cloudflare.F("**/blog/**"),
							Selector: cloudflare.F("article .post-body"),
						}}),
						IncludeHeaders: cloudflare.F(map[string]string{
							"foo": "string",
						}),
						IncludeImages:       cloudflare.F(true),
						SpecificSitemaps:    cloudflare.F([]string{"https://example.com/sitemap.xml", "https://example.com/blog-sitemap.xml"}),
						UseBrowserRendering: cloudflare.F(true),
					}),
					ParseType: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap),
					StoreOptions: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerStoreOptions{
						StorageID:      cloudflare.F("storage_id"),
						R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
						StorageType:    cloudflare.F(r2.ProviderR2),
					}),
				}),
			}),
			Summarization:                  cloudflare.F(true),
			SummarizationModel:             cloudflare.F(ai_search.InstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast),
			SyncInterval:                   cloudflare.F(ai_search.InstanceUpdateParamsSyncInterval900),
			SystemPromptAISearch:           cloudflare.F("system_prompt_ai_search"),
			SystemPromptIndexSummarization: cloudflare.F("system_prompt_index_summarization"),
			SystemPromptRewriteQuery:       cloudflare.F("system_prompt_rewrite_query"),
			TokenID:                        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.List(context.TODO(), ai_search.InstanceListParams{
		AccountID:        cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		Namespace:        cloudflare.F("namespace"),
		OrderBy:          cloudflare.F(ai_search.InstanceListParamsOrderByCreatedAt),
		OrderByDirection: cloudflare.F(ai_search.InstanceListParamsOrderByDirectionAsc),
		Page:             cloudflare.F(int64(1)),
		PerPage:          cloudflare.F(int64(20)),
		Search:           cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Delete(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceDeleteParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceChatCompletionsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.ChatCompletions(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceChatCompletionsParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			Messages: cloudflare.F([]ai_search.InstanceChatCompletionsParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.InstanceChatCompletionsParamsMessagesRoleSystem),
			}}),
			AISearchOptions: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptions{
				Cache: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Model:  cloudflare.F(ai_search.InstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast),
			Stream: cloudflare.F(true),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceRead(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Read(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceReadParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceSearchWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Search(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceSearchParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AISearchOptions: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptions{
				Cache: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.InstanceSearchParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Messages: cloudflare.F([]ai_search.InstanceSearchParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.InstanceSearchParamsMessagesRoleSystem),
			}}),
			Query: cloudflare.F("x"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStats(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Stats(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceStatsParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

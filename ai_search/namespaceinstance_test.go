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

func TestNamespaceInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.New(
		context.TODO(),
		"my-namespace",
		ai_search.NamespaceInstanceNewParams{
			AccountID:      cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			ID:             cloudflare.F("my-ai-search"),
			AIGatewayID:    cloudflare.F("ai_gateway_id"),
			AISearchModel:  cloudflare.F(ai_search.NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
			Cache:          cloudflare.F(true),
			CacheThreshold: cloudflare.F(ai_search.NamespaceInstanceNewParamsCacheThresholdSuperStrictMatch),
			Chunk:          cloudflare.F(true),
			ChunkOverlap:   cloudflare.F(int64(0)),
			ChunkSize:      cloudflare.F(int64(64)),
			CustomMetadata: cloudflare.F([]ai_search.NamespaceInstanceNewParamsCustomMetadata{{
				DataType:  cloudflare.F(ai_search.NamespaceInstanceNewParamsCustomMetadataDataTypeText),
				FieldName: cloudflare.F("x"),
			}}),
			EmbeddingModel:      cloudflare.F(ai_search.NamespaceInstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
			FusionMethod:        cloudflare.F(ai_search.NamespaceInstanceNewParamsFusionMethodMax),
			HybridSearchEnabled: cloudflare.F(true),
			IndexMethod: cloudflare.F(ai_search.NamespaceInstanceNewParamsIndexMethod{
				Keyword: cloudflare.F(true),
				Vector:  cloudflare.F(true),
			}),
			IndexingOptions: cloudflare.F(ai_search.NamespaceInstanceNewParamsIndexingOptions{
				KeywordTokenizer: cloudflare.F(ai_search.NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizerPorter),
			}),
			MaxNumResults: cloudflare.F(int64(1)),
			Metadata: cloudflare.F(ai_search.NamespaceInstanceNewParamsMetadata{
				CreatedFromAISearchWizard: cloudflare.F(true),
				WorkerDomain:              cloudflare.F("worker_domain"),
			}),
			PublicEndpointParams: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParams{
				AuthorizedHosts: cloudflare.F([]string{"string"}),
				ChatCompletionsEndpoint: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint{
					Disabled: cloudflare.F(true),
				}),
				Enabled: cloudflare.F(true),
				Mcp: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParamsMcp{
					Description: cloudflare.F("description"),
					Disabled:    cloudflare.F(true),
				}),
				RateLimit: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParamsRateLimit{
					PeriodMs:  cloudflare.F(int64(60000)),
					Requests:  cloudflare.F(int64(1)),
					Technique: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed),
				}),
				SearchEndpoint: cloudflare.F(ai_search.NamespaceInstanceNewParamsPublicEndpointParamsSearchEndpoint{
					Disabled: cloudflare.F(true),
				}),
			}),
			Reranking:      cloudflare.F(true),
			RerankingModel: cloudflare.F(ai_search.NamespaceInstanceNewParamsRerankingModelCfBaaiBgeRerankerBase),
			RetrievalOptions: cloudflare.F(ai_search.NamespaceInstanceNewParamsRetrievalOptions{
				BoostBy: cloudflare.F([]ai_search.NamespaceInstanceNewParamsRetrievalOptionsBoostBy{{
					Field:     cloudflare.F("timestamp"),
					Direction: cloudflare.F(ai_search.NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionDesc),
				}}),
				KeywordMatchMode: cloudflare.F(ai_search.NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchModeAnd),
			}),
			RewriteModel:   cloudflare.F(ai_search.NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
			RewriteQuery:   cloudflare.F(true),
			ScoreThreshold: cloudflare.F(0.000000),
			Source:         cloudflare.F("source"),
			SourceParams: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParams{
				ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
				IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
				Prefix:         cloudflare.F("prefix"),
				R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
				WebCrawler: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawler{
					CrawlOptions: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptions{
						Depth:                cloudflare.F(1.000000),
						IncludeExternalLinks: cloudflare.F(true),
						IncludeSubdomains:    cloudflare.F(true),
						MaxAge:               cloudflare.F(0.000000),
						Source:               cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceAll),
					}),
					ParseOptions: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions{
						ContentSelector: cloudflare.F([]ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector{{
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
					ParseType: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap),
					StoreOptions: cloudflare.F(ai_search.NamespaceInstanceNewParamsSourceParamsWebCrawlerStoreOptions{
						StorageID:      cloudflare.F("storage_id"),
						R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
						StorageType:    cloudflare.F(r2.ProviderR2),
					}),
				}),
			}),
			SyncInterval: cloudflare.F(ai_search.NamespaceInstanceNewParamsSyncInterval900),
			TokenID:      cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:         cloudflare.F(ai_search.NamespaceInstanceNewParamsTypeR2),
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

func TestNamespaceInstanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.Update(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceUpdateParams{
			AccountID:      cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AIGatewayID:    cloudflare.F("ai_gateway_id"),
			AISearchModel:  cloudflare.F(ai_search.NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
			Cache:          cloudflare.F(true),
			CacheThreshold: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsCacheThresholdSuperStrictMatch),
			Chunk:          cloudflare.F(true),
			ChunkOverlap:   cloudflare.F(int64(0)),
			ChunkSize:      cloudflare.F(int64(64)),
			CustomMetadata: cloudflare.F([]ai_search.NamespaceInstanceUpdateParamsCustomMetadata{{
				DataType:  cloudflare.F(ai_search.NamespaceInstanceUpdateParamsCustomMetadataDataTypeText),
				FieldName: cloudflare.F("x"),
			}}),
			EmbeddingModel: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
			FusionMethod:   cloudflare.F(ai_search.NamespaceInstanceUpdateParamsFusionMethodMax),
			IndexMethod: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsIndexMethod{
				Keyword: cloudflare.F(true),
				Vector:  cloudflare.F(true),
			}),
			IndexingOptions: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsIndexingOptions{
				KeywordTokenizer: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter),
			}),
			MaxNumResults: cloudflare.F(int64(1)),
			Metadata: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsMetadata{
				CreatedFromAISearchWizard: cloudflare.F(true),
				WorkerDomain:              cloudflare.F("worker_domain"),
			}),
			Paused: cloudflare.F(true),
			PublicEndpointParams: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParams{
				AuthorizedHosts: cloudflare.F([]string{"string"}),
				ChatCompletionsEndpoint: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint{
					Disabled: cloudflare.F(true),
				}),
				Enabled: cloudflare.F(true),
				Mcp: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParamsMcp{
					Description: cloudflare.F("description"),
					Disabled:    cloudflare.F(true),
				}),
				RateLimit: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimit{
					PeriodMs:  cloudflare.F(int64(60000)),
					Requests:  cloudflare.F(int64(1)),
					Technique: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed),
				}),
				SearchEndpoint: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsPublicEndpointParamsSearchEndpoint{
					Disabled: cloudflare.F(true),
				}),
			}),
			Reranking:      cloudflare.F(true),
			RerankingModel: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase),
			RetrievalOptions: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsRetrievalOptions{
				BoostBy: cloudflare.F([]ai_search.NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy{{
					Field:     cloudflare.F("timestamp"),
					Direction: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc),
				}}),
				KeywordMatchMode: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd),
			}),
			RewriteModel:   cloudflare.F(ai_search.NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
			RewriteQuery:   cloudflare.F(true),
			ScoreThreshold: cloudflare.F(0.000000),
			SourceParams: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParams{
				ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
				IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
				Prefix:         cloudflare.F("prefix"),
				R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
				WebCrawler: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawler{
					CrawlOptions: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptions{
						Depth:                cloudflare.F(1.000000),
						IncludeExternalLinks: cloudflare.F(true),
						IncludeSubdomains:    cloudflare.F(true),
						MaxAge:               cloudflare.F(0.000000),
						Source:               cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceAll),
					}),
					ParseOptions: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions{
						ContentSelector: cloudflare.F([]ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector{{
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
					ParseType: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap),
					StoreOptions: cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSourceParamsWebCrawlerStoreOptions{
						StorageID:      cloudflare.F("storage_id"),
						R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
						StorageType:    cloudflare.F(r2.ProviderR2),
					}),
				}),
			}),
			Summarization:                  cloudflare.F(true),
			SummarizationModel:             cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast),
			SyncInterval:                   cloudflare.F(ai_search.NamespaceInstanceUpdateParamsSyncInterval900),
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

func TestNamespaceInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.List(
		context.TODO(),
		"my-namespace",
		ai_search.NamespaceInstanceListParams{
			AccountID:        cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			Namespace:        cloudflare.F("namespace"),
			OrderBy:          cloudflare.F(ai_search.NamespaceInstanceListParamsOrderByCreatedAt),
			OrderByDirection: cloudflare.F(ai_search.NamespaceInstanceListParamsOrderByDirectionAsc),
			Page:             cloudflare.F(int64(1)),
			PerPage:          cloudflare.F(int64(20)),
			Search:           cloudflare.F("search"),
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

func TestNamespaceInstanceDelete(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.Delete(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceDeleteParams{
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

func TestNamespaceInstanceChatCompletionsWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.ChatCompletions(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceChatCompletionsParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			Messages: cloudflare.F([]ai_search.NamespaceInstanceChatCompletionsParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsMessagesRoleSystem),
			}}),
			AISearchOptions: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptions{
				Cache: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Model:  cloudflare.F(ai_search.NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast),
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

func TestNamespaceInstanceRead(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.Read(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceReadParams{
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

func TestNamespaceInstanceSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.Search(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceSearchParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AISearchOptions: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptions{
				Cache: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Messages: cloudflare.F([]ai_search.NamespaceInstanceSearchParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.NamespaceInstanceSearchParamsMessagesRoleSystem),
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

func TestNamespaceInstanceStats(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Instances.Stats(
		context.TODO(),
		"my-namespace",
		"my-ai-search",
		ai_search.NamespaceInstanceStatsParams{
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

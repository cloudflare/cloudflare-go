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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.New(context.TODO(), ai_search.InstanceNewParams{
		AccountID:     cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		ID:            cloudflare.F("my-ai-search"),
		Source:        cloudflare.F("source"),
		Type:          cloudflare.F(ai_search.InstanceNewParamsTypeR2),
		AIGatewayID:   cloudflare.F("ai_gateway_id"),
		AISearchModel: cloudflare.F(ai_search.InstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
		Chunk:         cloudflare.F(true),
		ChunkOverlap:  cloudflare.F(int64(0)),
		ChunkSize:     cloudflare.F(int64(64)),
		CustomMetadata: cloudflare.F([]ai_search.InstanceNewParamsCustomMetadata{{
			DataType:  cloudflare.F(ai_search.InstanceNewParamsCustomMetadataDataTypeText),
			FieldName: cloudflare.F("x"),
		}}),
		EmbeddingModel:      cloudflare.F(ai_search.InstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
		HybridSearchEnabled: cloudflare.F(true),
		MaxNumResults:       cloudflare.F(int64(1)),
		Metadata: cloudflare.F(ai_search.InstanceNewParamsMetadata{
			CreatedFromAISearchWizard: cloudflare.F(true),
			WorkerDomain:              cloudflare.F("worker_domain"),
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
		RewriteModel:   cloudflare.F(ai_search.InstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
		RewriteQuery:   cloudflare.F(true),
		ScoreThreshold: cloudflare.F(0.000000),
		SourceParams: cloudflare.F(ai_search.InstanceNewParamsSourceParams{
			ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
			IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
			Prefix:         cloudflare.F("prefix"),
			R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
			WebCrawler: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawler{
				ParseOptions: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerParseOptions{
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
		TokenID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			EmbeddingModel:      cloudflare.F(ai_search.InstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b),
			HybridSearchEnabled: cloudflare.F(true),
			MaxNumResults:       cloudflare.F(int64(1)),
			Metadata: cloudflare.F(ai_search.InstanceUpdateParamsMetadata{
				CreatedFromAISearchWizard: cloudflare.F(true),
				WorkerDomain:              cloudflare.F("worker_domain"),
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
			RewriteModel:   cloudflare.F(ai_search.InstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
			RewriteQuery:   cloudflare.F(true),
			ScoreThreshold: cloudflare.F(0.000000),
			SourceParams: cloudflare.F(ai_search.InstanceUpdateParamsSourceParams{
				ExcludeItems:   cloudflare.F([]string{"/admin/**", "/private/**", "**\\temp\\**"}),
				IncludeItems:   cloudflare.F([]string{"/blog/**", "/docs/**/*.html", "**\\blog\\**.html"}),
				Prefix:         cloudflare.F("prefix"),
				R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
				WebCrawler: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawler{
					ParseOptions: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerParseOptions{
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.List(context.TODO(), ai_search.InstanceListParams{
		AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
		Search:    cloudflare.F("search"),
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
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					MatchThreshold:  cloudflare.F(0.000000),
					MaxNumResults:   cloudflare.F(int64(1)),
					RetrievalType:   cloudflare.F(ai_search.InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure: cloudflare.F(true),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.AISearch.Instances.Search(
		context.TODO(),
		"my-ai-search",
		ai_search.InstanceSearchParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			Messages: cloudflare.F([]ai_search.InstanceSearchParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.InstanceSearchParamsMessagesRoleSystem),
			}}),
			AISearchOptions: cloudflare.F(ai_search.InstanceSearchParamsAISearchOptions{
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
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					MatchThreshold:  cloudflare.F(0.000000),
					MaxNumResults:   cloudflare.F(int64(1)),
					RetrievalType:   cloudflare.F(ai_search.InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure: cloudflare.F(true),
				}),
			}),
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

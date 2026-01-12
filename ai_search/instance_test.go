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
		AccountID:           cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		ID:                  cloudflare.F("my-ai-search"),
		Source:              cloudflare.F("source"),
		TokenID:             cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:                cloudflare.F(ai_search.InstanceNewParamsTypeR2),
		AIGatewayID:         cloudflare.F("ai_gateway_id"),
		AISearchModel:       cloudflare.F(ai_search.InstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
		Chunk:               cloudflare.F(true),
		ChunkOverlap:        cloudflare.F(int64(0)),
		ChunkSize:           cloudflare.F(int64(64)),
		EmbeddingModel:      cloudflare.F(ai_search.InstanceNewParamsEmbeddingModelCfBaaiBgeM3),
		HybridSearchEnabled: cloudflare.F(true),
		MaxNumResults:       cloudflare.F(int64(1)),
		Metadata: cloudflare.F(ai_search.InstanceNewParamsMetadata{
			CreatedFromAISearchWizard: cloudflare.F(true),
			WorkerDomain:              cloudflare.F("worker_domain"),
		}),
		PublicEndpointParams: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParams{
			AuthorizedHosts: cloudflare.F([]string{"string"}),
			Enabled:         cloudflare.F(true),
			RateLimit: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsRateLimit{
				PeriodMs:  cloudflare.F(int64(60000)),
				Requests:  cloudflare.F(int64(1)),
				Technique: cloudflare.F(ai_search.InstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed),
			}),
		}),
		Reranking:      cloudflare.F(true),
		RerankingModel: cloudflare.F(ai_search.InstanceNewParamsRerankingModelCfBaaiBgeRerankerBase),
		RewriteModel:   cloudflare.F(ai_search.InstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
		RewriteQuery:   cloudflare.F(true),
		ScoreThreshold: cloudflare.F(0.000000),
		SourceParams: cloudflare.F(ai_search.InstanceNewParamsSourceParams{
			ExcludeItems:   cloudflare.F([]string{"/admin/*", "/private/**", "*\\temp\\*"}),
			IncludeItems:   cloudflare.F([]string{"/blog/*", "/docs/**/*.html", "*\\blog\\*.html"}),
			Prefix:         cloudflare.F("prefix"),
			R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
			WebCrawler: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawler{
				ParseOptions: cloudflare.F(ai_search.InstanceNewParamsSourceParamsWebCrawlerParseOptions{
					IncludeHeaders: cloudflare.F(map[string]string{
						"foo": "string",
					}),
					IncludeImages:       cloudflare.F(true),
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
			AccountID:           cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AIGatewayID:         cloudflare.F("ai_gateway_id"),
			AISearchModel:       cloudflare.F(ai_search.InstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast),
			Cache:               cloudflare.F(true),
			CacheThreshold:      cloudflare.F(ai_search.InstanceUpdateParamsCacheThresholdSuperStrictMatch),
			Chunk:               cloudflare.F(true),
			ChunkOverlap:        cloudflare.F(int64(0)),
			ChunkSize:           cloudflare.F(int64(64)),
			EmbeddingModel:      cloudflare.F(ai_search.InstanceUpdateParamsEmbeddingModelCfBaaiBgeM3),
			HybridSearchEnabled: cloudflare.F(true),
			MaxNumResults:       cloudflare.F(int64(1)),
			Metadata: cloudflare.F(ai_search.InstanceUpdateParamsMetadata{
				CreatedFromAISearchWizard: cloudflare.F(true),
				WorkerDomain:              cloudflare.F("worker_domain"),
			}),
			Paused: cloudflare.F(true),
			PublicEndpointParams: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParams{
				AuthorizedHosts: cloudflare.F([]string{"string"}),
				Enabled:         cloudflare.F(true),
				RateLimit: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsRateLimit{
					PeriodMs:  cloudflare.F(int64(60000)),
					Requests:  cloudflare.F(int64(1)),
					Technique: cloudflare.F(ai_search.InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed),
				}),
			}),
			Reranking:      cloudflare.F(true),
			RerankingModel: cloudflare.F(ai_search.InstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase),
			RewriteModel:   cloudflare.F(ai_search.InstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
			RewriteQuery:   cloudflare.F(true),
			ScoreThreshold: cloudflare.F(0.000000),
			SourceParams: cloudflare.F(ai_search.InstanceUpdateParamsSourceParams{
				ExcludeItems:   cloudflare.F([]string{"/admin/*", "/private/**", "*\\temp\\*"}),
				IncludeItems:   cloudflare.F([]string{"/blog/*", "/docs/**/*.html", "*\\blog\\*.html"}),
				Prefix:         cloudflare.F("prefix"),
				R2Jurisdiction: cloudflare.F("r2_jurisdiction"),
				WebCrawler: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawler{
					ParseOptions: cloudflare.F(ai_search.InstanceUpdateParamsSourceParamsWebCrawlerParseOptions{
						IncludeHeaders: cloudflare.F(map[string]string{
							"foo": "string",
						}),
						IncludeImages:       cloudflare.F(true),
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

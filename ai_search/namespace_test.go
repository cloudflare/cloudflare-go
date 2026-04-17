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
)

func TestNamespaceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.New(context.TODO(), ai_search.NamespaceNewParams{
		AccountID:   cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		Name:        cloudflare.F("name"),
		Description: cloudflare.F("Production environment"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Update(
		context.TODO(),
		"production",
		ai_search.NamespaceUpdateParams{
			AccountID:   cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			Description: cloudflare.F("Production environment"),
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

func TestNamespaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.List(context.TODO(), ai_search.NamespaceListParams{
		AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(20)),
		Search:    cloudflare.F("prod"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNamespaceDelete(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Delete(
		context.TODO(),
		"production",
		ai_search.NamespaceDeleteParams{
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

func TestNamespaceChatCompletionsWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.ChatCompletions(
		context.TODO(),
		"my-namespace",
		ai_search.NamespaceChatCompletionsParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AISearchOptions: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptions{
				InstanceIDs: cloudflare.F([]string{"my-ai-search"}),
				Cache: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Messages: cloudflare.F([]ai_search.NamespaceChatCompletionsParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.NamespaceChatCompletionsParamsMessagesRoleSystem),
			}}),
			Model:  cloudflare.F(ai_search.NamespaceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast),
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

func TestNamespaceRead(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Read(
		context.TODO(),
		"production",
		ai_search.NamespaceReadParams{
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

func TestNamespaceSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.AISearch.Namespaces.Search(
		context.TODO(),
		"my-namespace",
		ai_search.NamespaceSearchParams{
			AccountID: cloudflare.F("c3dc5f0b34a14ff8e1b3ec04895e1b22"),
			AISearchOptions: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptions{
				InstanceIDs: cloudflare.F([]string{"my-ai-search"}),
				Cache: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsCache{
					CacheThreshold: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch),
					Enabled:        cloudflare.F(true),
				}),
				QueryRewrite: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsQueryRewrite{
					Enabled:       cloudflare.F(true),
					Model:         cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast),
					RewritePrompt: cloudflare.F("rewrite_prompt"),
				}),
				Reranking: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsReranking{
					Enabled:        cloudflare.F(true),
					MatchThreshold: cloudflare.F(0.000000),
					Model:          cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase),
				}),
				Retrieval: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRetrieval{
					BoostBy: cloudflare.F([]ai_search.NamespaceSearchParamsAISearchOptionsRetrievalBoostBy{{
						Field:     cloudflare.F("timestamp"),
						Direction: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc),
					}}),
					ContextExpansion: cloudflare.F(int64(0)),
					Filters: cloudflare.F(map[string]interface{}{
						"foo": "bar",
					}),
					FusionMethod:     cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRetrievalFusionMethodMax),
					KeywordMatchMode: cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd),
					MatchThreshold:   cloudflare.F(0.000000),
					MaxNumResults:    cloudflare.F(int64(1)),
					RetrievalType:    cloudflare.F(ai_search.NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector),
					ReturnOnFailure:  cloudflare.F(true),
				}),
			}),
			Messages: cloudflare.F([]ai_search.NamespaceSearchParamsMessage{{
				Content: cloudflare.F("content"),
				Role:    cloudflare.F(ai_search.NamespaceSearchParamsMessagesRoleSystem),
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

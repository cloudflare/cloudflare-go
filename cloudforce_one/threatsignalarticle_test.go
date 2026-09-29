// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestThreatSignalArticleListWithOptionalParams(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatSignals.Articles.List(context.TODO(), cloudforce_one.ThreatSignalArticleListParams{
		AccountID:       cloudflare.F("account_id"),
		ArticleID:       cloudflare.F([]string{"550e8400-e29b-41d4-a716-446655440000", "660e8400-e29b-41d4-a716-446655440000"}),
		Cursor:          cloudflare.F("x"),
		FeedCategory:    cloudflare.F("feed_category"),
		FeedID:          cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FetchedAfter:    cloudflare.F(time.Now()),
		FetchedBefore:   cloudflare.F(time.Now()),
		IncludeTotal:    cloudflare.F(true),
		PerPage:         cloudflare.F(int64(1)),
		PublishedAfter:  cloudflare.F(time.Now()),
		PublishedBefore: cloudflare.F(time.Now()),
		Read:            cloudflare.F(true),
		Search:          cloudflare.F("x"),
		Sort:            cloudflare.F("sort"),
		SourceType:      cloudflare.F(cloudforce_one.ThreatSignalArticleListParamsSourceTypeCurated),
		Tag:             cloudflare.F("tag"),
		TagAppliedBy:    cloudflare.F(cloudforce_one.ThreatSignalArticleListParamsTagAppliedByAI),
		TagCategory:     cloudflare.F("tag_category"),
		TagCategoryID:   cloudflare.F([]string{"660e8400-e29b-41d4-a716-446655440000"}),
		TagID:           cloudflare.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatSignalArticleBulkEdit(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatSignals.Articles.BulkEdit(context.TODO(), cloudforce_one.ThreatSignalArticleBulkEditParams{
		AccountID:  cloudflare.F("account_id"),
		ArticleIDs: cloudflare.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Read:       cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestThreatSignalArticleEdit(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatSignals.Articles.Edit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cloudforce_one.ThreatSignalArticleEditParams{
			AccountID: cloudflare.F("account_id"),
			Read:      cloudflare.F(true),
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

func TestThreatSignalArticleGet(t *testing.T) {
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
	_, err := client.CloudforceOne.ThreatSignals.Articles.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cloudforce_one.ThreatSignalArticleGetParams{
			AccountID: cloudflare.F("account_id"),
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

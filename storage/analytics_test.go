// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/storage"
)

func TestAnalyticsListWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Analytics.List(context.TODO(), storage.AnalyticsListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(storage.AnalyticsListParamsQuery{
			Dimensions: cloudflare.F([]storage.AnalyticsListParamsQueryDimension{storage.AnalyticsListParamsQueryDimensionAccountID, storage.AnalyticsListParamsQueryDimensionResponseCode, storage.AnalyticsListParamsQueryDimensionRequestType}),
			Filters:    cloudflare.F("requestType==read AND responseCode!=200"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]storage.AnalyticsListParamsQueryMetric{storage.AnalyticsListParamsQueryMetricRequests, storage.AnalyticsListParamsQueryMetricWriteKiB, storage.AnalyticsListParamsQueryMetricReadKiB}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+requests", "-responseCode"}),
			Until:      cloudflare.F(time.Now()),
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

func TestAnalyticsStoredWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Analytics.Stored(context.TODO(), storage.AnalyticsStoredParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(storage.AnalyticsStoredParamsQuery{
			Dimensions: cloudflare.F([]storage.AnalyticsStoredParamsQueryDimension{storage.AnalyticsStoredParamsQueryDimensionNamespaceID}),
			Filters:    cloudflare.F("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]storage.AnalyticsStoredParamsQueryMetric{storage.AnalyticsStoredParamsQueryMetricStoredBytes, storage.AnalyticsStoredParamsQueryMetricStoredKeys}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+storedBytes", "-namespaceId"}),
			Until:      cloudflare.F(time.Now()),
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

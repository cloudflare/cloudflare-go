// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestStorageAnalyticsListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Storage.Analytics.List(context.TODO(), cloudflare.StorageAnalyticsListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(cloudflare.StorageAnalyticsListParamsQuery{
			Dimensions: cloudflare.F([]cloudflare.StorageAnalyticsListParamsQueryDimension{cloudflare.StorageAnalyticsListParamsQueryDimensionAccountID, cloudflare.StorageAnalyticsListParamsQueryDimensionResponseCode, cloudflare.StorageAnalyticsListParamsQueryDimensionRequestType}),
			Filters:    cloudflare.F("requestType==read AND responseCode!=200"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]cloudflare.StorageAnalyticsListParamsQueryMetric{cloudflare.StorageAnalyticsListParamsQueryMetricRequests, cloudflare.StorageAnalyticsListParamsQueryMetricWriteKiB, cloudflare.StorageAnalyticsListParamsQueryMetricReadKiB}),
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

func TestStorageAnalyticsStoredWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Storage.Analytics.Stored(context.TODO(), cloudflare.StorageAnalyticsStoredParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(cloudflare.StorageAnalyticsStoredParamsQuery{
			Dimensions: cloudflare.F([]cloudflare.StorageAnalyticsStoredParamsQueryDimension{cloudflare.StorageAnalyticsStoredParamsQueryDimensionNamespaceID}),
			Filters:    cloudflare.F("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]cloudflare.StorageAnalyticsStoredParamsQueryMetric{cloudflare.StorageAnalyticsStoredParamsQueryMetricStoredBytes, cloudflare.StorageAnalyticsStoredParamsQueryMetricStoredKeys}),
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

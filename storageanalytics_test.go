// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Storage.Analytics.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.StorageAnalyticsListParams{
			Query: cloudflare.F(cloudflare.StorageAnalyticsListParamsQuery{
				Dimensions: cloudflare.F([]cloudflare.StorageAnalyticsListParamsQueryDimension{cloudflare.StorageAnalyticsListParamsQueryDimensionAccountID, cloudflare.StorageAnalyticsListParamsQueryDimensionResponseCode, cloudflare.StorageAnalyticsListParamsQueryDimensionRequestType}),
				Filters:    cloudflare.F("requestType==read AND responseCode!=200"),
				Limit:      cloudflare.F(int64(0)),
				Metrics:    cloudflare.F([]cloudflare.StorageAnalyticsListParamsQueryMetric{cloudflare.StorageAnalyticsListParamsQueryMetricRequests, cloudflare.StorageAnalyticsListParamsQueryMetricWriteKiB, cloudflare.StorageAnalyticsListParamsQueryMetricReadKiB}),
				Since:      cloudflare.F(time.Now()),
				Sort:       cloudflare.F([]interface{}{"+requests", "-responseCode"}),
				Until:      cloudflare.F(time.Now()),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Storage.Analytics.Stored(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.StorageAnalyticsStoredParams{
			Query: cloudflare.F(cloudflare.StorageAnalyticsStoredParamsQuery{
				Dimensions: cloudflare.F([]cloudflare.StorageAnalyticsStoredParamsQueryDimension{cloudflare.StorageAnalyticsStoredParamsQueryDimensionNamespaceID}),
				Filters:    cloudflare.F("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
				Limit:      cloudflare.F(int64(0)),
				Metrics:    cloudflare.F([]cloudflare.StorageAnalyticsStoredParamsQueryMetric{cloudflare.StorageAnalyticsStoredParamsQueryMetricStoredBytes, cloudflare.StorageAnalyticsStoredParamsQueryMetricStoredKeys}),
				Since:      cloudflare.F(time.Now()),
				Sort:       cloudflare.F([]interface{}{"+storedBytes", "-namespaceId"}),
				Until:      cloudflare.F(time.Now()),
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

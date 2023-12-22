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

func TestAccountStorageAnalyticsListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Storages.Analytics.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStorageAnalyticsListParams{
			Query: cloudflare.F(cloudflare.AccountStorageAnalyticsListParamsQuery{
				Dimensions: cloudflare.F([]cloudflare.AccountStorageAnalyticsListParamsQueryDimension{cloudflare.AccountStorageAnalyticsListParamsQueryDimensionAccountID, cloudflare.AccountStorageAnalyticsListParamsQueryDimensionResponseCode}),
				Filters:    cloudflare.F[any]("requestType==read AND responseCode!=200"),
				Limit:      cloudflare.F(int64(0)),
				Metrics:    cloudflare.F([]cloudflare.AccountStorageAnalyticsListParamsQueryMetric{cloudflare.AccountStorageAnalyticsListParamsQueryMetricRequests, cloudflare.AccountStorageAnalyticsListParamsQueryMetricReadKiB}),
				Since:      cloudflare.F(time.Now()),
				Sort: cloudflare.F[any](map[string]interface{}{
					"0": "+requests",
					"1": "-responseCode",
				}),
				Until: cloudflare.F(time.Now()),
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

func TestAccountStorageAnalyticsStoredWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Accounts.Storages.Analytics.Stored(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStorageAnalyticsStoredParams{
			Query: cloudflare.F(cloudflare.AccountStorageAnalyticsStoredParamsQuery{
				Dimensions: cloudflare.F([]cloudflare.AccountStorageAnalyticsStoredParamsQueryDimension{cloudflare.AccountStorageAnalyticsStoredParamsQueryDimensionNamespaceID}),
				Filters:    cloudflare.F[any]("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
				Limit:      cloudflare.F(int64(0)),
				Metrics:    cloudflare.F([]cloudflare.AccountStorageAnalyticsStoredParamsQueryMetric{cloudflare.AccountStorageAnalyticsStoredParamsQueryMetricStoredBytes, cloudflare.AccountStorageAnalyticsStoredParamsQueryMetricStoredKeys}),
				Since:      cloudflare.F(time.Now()),
				Sort: cloudflare.F[any](map[string]interface{}{
					"0": "+storedBytes",
					"1": "-namespaceId",
				}),
				Until: cloudflare.F(time.Now()),
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

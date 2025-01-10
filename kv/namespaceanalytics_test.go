// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/kv"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestNamespaceAnalyticsListWithOptionalParams(t *testing.T) {
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
	_, err := client.KV.Namespaces.Analytics.List(context.TODO(), kv.NamespaceAnalyticsListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(kv.NamespaceAnalyticsListParamsQuery{
			Dimensions: cloudflare.F([]kv.NamespaceAnalyticsListParamsQueryDimension{kv.NamespaceAnalyticsListParamsQueryDimensionAccountID}),
			Filters:    cloudflare.F("requestType==read AND responseCode!=200"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]kv.NamespaceAnalyticsListParamsQueryMetric{kv.NamespaceAnalyticsListParamsQueryMetricRequests}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]string{"+requests", "-responseCode"}),
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

func TestNamespaceAnalyticsStoredWithOptionalParams(t *testing.T) {
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
	_, err := client.KV.Namespaces.Analytics.Stored(context.TODO(), kv.NamespaceAnalyticsStoredParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: cloudflare.F(kv.NamespaceAnalyticsStoredParamsQuery{
			Dimensions: cloudflare.F([]kv.NamespaceAnalyticsStoredParamsQueryDimension{kv.NamespaceAnalyticsStoredParamsQueryDimensionNamespaceID}),
			Filters:    cloudflare.F("namespaceId==a4e8cbb7-1b58-4990-925e-e026d40c4c64"),
			Limit:      cloudflare.F(int64(0)),
			Metrics:    cloudflare.F([]kv.NamespaceAnalyticsStoredParamsQueryMetric{kv.NamespaceAnalyticsStoredParamsQueryMetricStoredBytes}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]string{"+storedBytes", "-namespaceId"}),
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

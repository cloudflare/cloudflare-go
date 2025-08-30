// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/brand_protection"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestQueryNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test, 401 Unauthorized")
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
	err := client.BrandProtection.Queries.New(context.TODO(), brand_protection.QueryNewParams{
		AccountID:     cloudflare.F("x"),
		ID:            cloudflare.F("id"),
		QueryScan:     cloudflare.F(true),
		QueryTag:      cloudflare.F("tag"),
		MaxTime:       cloudflare.F(time.Now()),
		MinTime:       cloudflare.F(time.Now()),
		BodyScan:      cloudflare.F(true),
		StringMatches: cloudflare.F[any](map[string]interface{}{}),
		BodyTag:       cloudflare.F("tag"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQueryDeleteWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test, 401 Unauthorized")
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
	err := client.BrandProtection.Queries.Delete(context.TODO(), brand_protection.QueryDeleteParams{
		AccountID: cloudflare.F("x"),
		ID:        cloudflare.F("id"),
		Scan:      cloudflare.F(true),
		Tag:       cloudflare.F("tag"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQueryBulkWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test, 401 Unauthorized")
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
	err := client.BrandProtection.Queries.Bulk(context.TODO(), brand_protection.QueryBulkParams{
		AccountID: cloudflare.F("x"),
		Queries: cloudflare.F([]map[string]interface{}{{
			"foo": "bar",
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

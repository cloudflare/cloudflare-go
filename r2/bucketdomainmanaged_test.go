// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/r2"
)

func TestBucketDomainManagedUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.R2.Buckets.Domains.Managed.Update(
		context.TODO(),
		"example-bucket",
		r2.BucketDomainManagedUpdateParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Enabled:          cloudflare.F(true),
			CfR2Jurisdiction: cloudflare.F(r2.BucketDomainManagedUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketDomainManagedListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.R2.Buckets.Domains.Managed.List(
		context.TODO(),
		"example-bucket",
		r2.BucketDomainManagedListParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketDomainManagedListParamsCfR2JurisdictionDefault),
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

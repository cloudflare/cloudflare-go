// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/r2"
)

func TestBucketSippyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.Sippy.Update(
		context.TODO(),
		"example-bucket",
		r2.BucketSippyUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: r2.BucketSippyUpdateParamsBodyR2EnableSippyAws{
				Destination: cloudflare.F(r2.BucketSippyUpdateParamsBodyR2EnableSippyAwsDestination{
					AccessKeyID:     cloudflare.F("accessKeyId"),
					Provider:        cloudflare.F(r2.ProviderR2),
					SecretAccessKey: cloudflare.F("secretAccessKey"),
				}),
				Source: cloudflare.F(r2.BucketSippyUpdateParamsBodyR2EnableSippyAwsSource{
					AccessKeyID:     cloudflare.F("accessKeyId"),
					Bucket:          cloudflare.F("bucket"),
					Provider:        cloudflare.F(r2.BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProviderAws),
					Region:          cloudflare.F("region"),
					SecretAccessKey: cloudflare.F("secretAccessKey"),
				}),
			},
			CfR2Jurisdiction: cloudflare.F(r2.BucketSippyUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketSippyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.Sippy.Delete(
		context.TODO(),
		"example-bucket",
		r2.BucketSippyDeleteParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketSippyDeleteParamsCfR2JurisdictionDefault),
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

func TestBucketSippyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.Sippy.Get(
		context.TODO(),
		"example-bucket",
		r2.BucketSippyGetParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketSippyGetParamsCfR2JurisdictionDefault),
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

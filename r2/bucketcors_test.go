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

func TestBucketCORSUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.CORS.Update(
		context.TODO(),
		"example-bucket",
		r2.BucketCORSUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: cloudflare.F([]r2.BucketCORSUpdateParamsRule{{
				Allowed: cloudflare.F(r2.BucketCORSUpdateParamsRulesAllowed{
					Methods: cloudflare.F([]r2.BucketCORSUpdateParamsRulesAllowedMethod{r2.BucketCORSUpdateParamsRulesAllowedMethodGet}),
					Origins: cloudflare.F([]string{"http://localhost:3000"}),
					Headers: cloudflare.F([]string{"x-requested-by"}),
				}),
				ID:            cloudflare.F("Allow Local Development"),
				ExposeHeaders: cloudflare.F([]string{"Content-Encoding"}),
				MaxAgeSeconds: cloudflare.F(3600.000000),
			}}),
			CfR2Jurisdiction: cloudflare.F(r2.BucketCORSUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketCORSDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.CORS.Delete(
		context.TODO(),
		"example-bucket",
		r2.BucketCORSDeleteParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketCORSDeleteParamsCfR2JurisdictionDefault),
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

func TestBucketCORSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.CORS.Get(
		context.TODO(),
		"example-bucket",
		r2.BucketCORSGetParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketCORSGetParamsCfR2JurisdictionDefault),
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

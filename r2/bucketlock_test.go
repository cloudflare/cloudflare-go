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

func TestBucketLockUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate auth errors on test suite")
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
	_, err := client.R2.Buckets.Locks.Update(
		context.TODO(),
		"example-bucket",
		r2.BucketLockUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: cloudflare.F([]r2.BucketLockUpdateParamsRule{{
				ID: cloudflare.F("Lock all objects for 24 hours"),
				Condition: cloudflare.F[r2.BucketLockUpdateParamsRulesConditionUnion](r2.BucketLockUpdateParamsRulesConditionR2LockRuleAgeCondition{
					MaxAgeSeconds: cloudflare.F(int64(100)),
					Type:          cloudflare.F(r2.BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionTypeAge),
				}),
				Enabled: cloudflare.F(true),
				Prefix:  cloudflare.F("prefix"),
			}}),
			Jurisdiction: cloudflare.F(r2.BucketLockUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketLockGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate auth errors on test suite")
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
	_, err := client.R2.Buckets.Locks.Get(
		context.TODO(),
		"example-bucket",
		r2.BucketLockGetParams{
			AccountID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Jurisdiction: cloudflare.F(r2.BucketLockGetParamsCfR2JurisdictionDefault),
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

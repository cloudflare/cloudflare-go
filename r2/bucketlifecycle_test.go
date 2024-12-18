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

func TestBucketLifecycleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.Lifecycle.Update(
		context.TODO(),
		"example-bucket",
		r2.BucketLifecycleUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: cloudflare.F([]r2.BucketLifecycleUpdateParamsRule{{
				ID: cloudflare.F("Expire all objects older than 24 hours"),
				Conditions: cloudflare.F(r2.BucketLifecycleUpdateParamsRulesConditions{
					Prefix: cloudflare.F("prefix"),
				}),
				Enabled: cloudflare.F(true),
				AbortMultipartUploadsTransition: cloudflare.F(r2.BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransition{
					Condition: cloudflare.F(r2.BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionCondition{
						MaxAge: cloudflare.F(int64(0)),
						Type:   cloudflare.F(r2.BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionTypeAge),
					}),
				}),
				DeleteObjectsTransition: cloudflare.F(r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransition{
					Condition: cloudflare.F[r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion](r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition{
						MaxAge: cloudflare.F(int64(0)),
						Type:   cloudflare.F(r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge),
					}),
				}),
				StorageClassTransitions: cloudflare.F([]r2.BucketLifecycleUpdateParamsRulesStorageClassTransition{{
					Condition: cloudflare.F[r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion](r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeCondition{
						MaxAge: cloudflare.F(int64(0)),
						Type:   cloudflare.F(r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge),
					}),
					StorageClass: cloudflare.F(r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClassInfrequentAccess),
				}}),
			}}),
			CfR2Jurisdiction: cloudflare.F(r2.BucketLifecycleUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketLifecycleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.Lifecycle.Get(
		context.TODO(),
		"example-bucket",
		r2.BucketLifecycleGetParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketLifecycleGetParamsCfR2JurisdictionDefault),
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

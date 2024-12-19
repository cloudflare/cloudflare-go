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

func TestBucketEventNotificationConfigurationQueueUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.EventNotifications.Configuration.Queues.Update(
		context.TODO(),
		"example-bucket",
		"queue_id",
		r2.BucketEventNotificationConfigurationQueueUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: cloudflare.F([]r2.BucketEventNotificationConfigurationQueueUpdateParamsRule{{
				Actions:     cloudflare.F([]r2.BucketEventNotificationConfigurationQueueUpdateParamsRulesAction{r2.BucketEventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, r2.BucketEventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Description: cloudflare.F("Notifications from source bucket to queue"),
				Prefix:      cloudflare.F("img/"),
				Suffix:      cloudflare.F(".jpeg"),
			}}),
			CfR2Jurisdiction: cloudflare.F(r2.BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault),
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

func TestBucketEventNotificationConfigurationQueueDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.R2.Buckets.EventNotifications.Configuration.Queues.Delete(
		context.TODO(),
		"example-bucket",
		"queue_id",
		r2.BucketEventNotificationConfigurationQueueDeleteParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault),
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

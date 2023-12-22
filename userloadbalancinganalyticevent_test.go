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

func TestUserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.LoadBalancingAnalytics.Events.LoadBalancerHealthcheckEventsListHealthcheckEvents(context.TODO(), cloudflare.UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams{
		Identifier:    cloudflare.F[any]("17b5962d775c646f3f9725cbc7a53df4"),
		OriginHealthy: cloudflare.F(true),
		OriginName:    cloudflare.F("primary-dc-1"),
		PoolHealthy:   cloudflare.F(true),
		PoolName:      cloudflare.F("primary-dc"),
		Since:         cloudflare.F(time.Now()),
		Until:         cloudflare.F(time.Now()),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

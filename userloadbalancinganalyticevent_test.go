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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Users.LoadBalancingAnalytics.Events.LoadBalancerHealthcheckEventsListHealthcheckEvents(context.TODO(), cloudflare.UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams{
		OriginHealthy: cloudflare.F(true),
		OriginName:    cloudflare.F("primary-dc-1"),
		PoolHealthy:   cloudflare.F(true),
		PoolID:        cloudflare.F("17b5962d775c646f3f9725cbc7a53df4"),
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

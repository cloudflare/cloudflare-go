// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestRobotsTXTTopUserAgentDirectiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.RobotsTXT.Top.UserAgents.Directive(context.TODO(), radar.RobotsTXTTopUserAgentDirectiveParams{
		Date:              cloudflare.F([]time.Time{time.Now()}),
		Directive:         cloudflare.F(radar.RobotsTXTTopUserAgentDirectiveParamsDirectiveAllow),
		DomainCategory:    cloudflare.F([]string{"string"}),
		Format:            cloudflare.F(radar.RobotsTXTTopUserAgentDirectiveParamsFormatJson),
		Limit:             cloudflare.F(int64(5)),
		Name:              cloudflare.F([]string{"string"}),
		UserAgentCategory: cloudflare.F(radar.RobotsTXTTopUserAgentDirectiveParamsUserAgentCategoryAI),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

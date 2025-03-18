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

func TestLeakedCredentialSummaryBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.LeakedCredentials.Summary.BotClass(context.TODO(), radar.LeakedCredentialSummaryBotClassParams{
		Compromised: cloudflare.F([]radar.LeakedCredentialSummaryBotClassParamsCompromised{radar.LeakedCredentialSummaryBotClassParamsCompromisedClean}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.LeakedCredentialSummaryBotClassParamsFormatJson),
		Name:        cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLeakedCredentialSummaryCompromisedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.LeakedCredentials.Summary.Compromised(context.TODO(), radar.LeakedCredentialSummaryCompromisedParams{
		BotClass:  cloudflare.F([]radar.LeakedCredentialSummaryCompromisedParamsBotClass{radar.LeakedCredentialSummaryCompromisedParamsBotClassLikelyAutomated}),
		DateEnd:   cloudflare.F([]time.Time{time.Now()}),
		DateRange: cloudflare.F([]string{"7d"}),
		DateStart: cloudflare.F([]time.Time{time.Now()}),
		Format:    cloudflare.F(radar.LeakedCredentialSummaryCompromisedParamsFormatJson),
		Name:      cloudflare.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

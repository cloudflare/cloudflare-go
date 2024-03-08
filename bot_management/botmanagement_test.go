// File generated from our OpenAPI spec by Stainless.

package bot_management_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/bot_management"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestBotManagementUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BotManagement.Update(context.TODO(), bot_management.BotManagementUpdateParams{
		ZoneID:                       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AutoUpdateModel:              cloudflare.F(true),
		EnableJs:                     cloudflare.F(true),
		FightMode:                    cloudflare.F(true),
		OptimizeWordpress:            cloudflare.F(true),
		SbfmDefinitelyAutomated:      cloudflare.F(bot_management.BotManagementUpdateParamsSbfmDefinitelyAutomatedAllow),
		SbfmLikelyAutomated:          cloudflare.F(bot_management.BotManagementUpdateParamsSbfmLikelyAutomatedAllow),
		SbfmStaticResourceProtection: cloudflare.F(true),
		SbfmVerifiedBots:             cloudflare.F(bot_management.BotManagementUpdateParamsSbfmVerifiedBotsAllow),
		SuppressSessionScore:         cloudflare.F(false),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBotManagementGet(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.BotManagement.Get(context.TODO(), bot_management.BotManagementGetParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

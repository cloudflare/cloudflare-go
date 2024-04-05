// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
)

func TestRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleNewParams{
			ZoneID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:      cloudflare.F(waiting_rooms.RuleNewParamsActionBypassWaitingRoom),
			Expression:  cloudflare.F("ip.src in {10.20.30.40}"),
			Description: cloudflare.F("allow all traffic from 10.20.30.40"),
			Enabled:     cloudflare.F(true),
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

func TestRuleUpdate(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: cloudflare.F([]waiting_rooms.RuleParam{{
				Action:      cloudflare.F(waiting_rooms.RuleActionBypassWaitingRoom),
				Description: cloudflare.F("allow all traffic from 10.20.30.40"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src in {10.20.30.40}"),
			}, {
				Action:      cloudflare.F(waiting_rooms.RuleActionBypassWaitingRoom),
				Description: cloudflare.F("allow all traffic from 10.20.30.40"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src in {10.20.30.40}"),
			}, {
				Action:      cloudflare.F(waiting_rooms.RuleActionBypassWaitingRoom),
				Description: cloudflare.F("allow all traffic from 10.20.30.40"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src in {10.20.30.40}"),
			}}),
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

func TestRuleList(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.List(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.RuleListParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestRuleDelete(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.RuleDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:   cloudflare.F[any](map[string]interface{}{}),
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

func TestRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.WaitingRooms.Rules.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		waiting_rooms.RuleEditParams{
			ZoneID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:      cloudflare.F(waiting_rooms.RuleEditParamsActionBypassWaitingRoom),
			Expression:  cloudflare.F("ip.src in {10.20.30.40}"),
			Description: cloudflare.F("allow all traffic from 10.20.30.40"),
			Enabled:     cloudflare.F(true),
			Position: cloudflare.F[waiting_rooms.RuleEditParamsPositionUnion](waiting_rooms.RuleEditParamsPositionObject{
				Index: cloudflare.F(int64(0)),
			}),
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

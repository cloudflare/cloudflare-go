// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestDEXCommandNew(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.Commands.New(context.TODO(), zero_trust.DEXCommandNewParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Commands: cloudflare.F([]zero_trust.DEXCommandNewParamsCommand{{
			CommandType: cloudflare.F(zero_trust.DEXCommandNewParamsCommandsCommandTypePCAP),
			DeviceID:    cloudflare.F("device_id"),
			UserEmail:   cloudflare.F("user_email"),
			CommandArgs: cloudflare.F(zero_trust.DEXCommandNewParamsCommandsCommandArgs{
				Interfaces:      cloudflare.F([]zero_trust.DEXCommandNewParamsCommandsCommandArgsInterface{zero_trust.DEXCommandNewParamsCommandsCommandArgsInterfaceDefault}),
				MaxFileSizeMB:   cloudflare.F(1.000000),
				PacketSizeBytes: cloudflare.F(1.000000),
				TestAllRoutes:   cloudflare.F(true),
				TimeLimitMin:    cloudflare.F(1.000000),
			}),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDEXCommandListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.Commands.List(context.TODO(), zero_trust.DEXCommandListParams{
		AccountID:   cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Page:        cloudflare.F(1.000000),
		PerPage:     cloudflare.F(50.000000),
		CommandType: cloudflare.F("command_type"),
		DeviceID:    cloudflare.F("device_id"),
		From:        cloudflare.F(time.Now()),
		Status:      cloudflare.F(zero_trust.DEXCommandListParamsStatusPendingExec),
		To:          cloudflare.F(time.Now()),
		UserEmail:   cloudflare.F("user_email"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

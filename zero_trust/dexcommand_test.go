// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.DEX.Commands.New(context.TODO(), zero_trust.DEXCommandNewParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Commands: cloudflare.F([]zero_trust.DEXCommandNewParamsCommand{{
			DeviceID:  cloudflare.F("device_id"),
			Type:      cloudflare.F(zero_trust.DEXCommandNewParamsCommandsTypePCAP),
			UserEmail: cloudflare.F("user_email"),
			Args: cloudflare.F[zero_trust.DEXCommandNewParamsCommandsArgsUnion](zero_trust.DEXCommandNewParamsCommandsArgsWARPDiagArgs{
				TestAllRoutes: cloudflare.F(true),
			}),
			RegistrationID: cloudflare.F("registration_id"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.DEX.Commands.List(context.TODO(), zero_trust.DEXCommandListParams{
		AccountID:   cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		Page:        cloudflare.F(1.000000),
		PerPage:     cloudflare.F(10.000000),
		CommandType: cloudflare.F(zero_trust.DEXCommandListParamsCommandTypePCAP),
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
)

func TestDLPProfilePredefinedNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.New(context.TODO(), zero_trust.DLPProfilePredefinedNewParams{
		AccountID:           cloudflare.F("account_id"),
		ProfileID:           cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AIContextEnabled:    cloudflare.F(true),
		AllowedMatchCount:   cloudflare.F(int64(5)),
		ConfidenceThreshold: cloudflare.F("confidence_threshold"),
		ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
			Enabled: cloudflare.F(true),
			Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
				Files: cloudflare.F(true),
			}),
		}),
		Entries: cloudflare.F([]zero_trust.DLPProfilePredefinedNewParamsEntry{{
			ID:      cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Enabled: cloudflare.F(true),
		}}),
		OCREnabled: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDLPProfilePredefinedUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfilePredefinedUpdateParams{
			AccountID:           cloudflare.F("account_id"),
			AIContextEnabled:    cloudflare.F(true),
			AllowedMatchCount:   cloudflare.F(int64(5)),
			ConfidenceThreshold: cloudflare.F("confidence_threshold"),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Entries: cloudflare.F([]zero_trust.DLPProfilePredefinedUpdateParamsEntry{{
				ID:      cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Enabled: cloudflare.F(true),
			}}),
			OCREnabled: cloudflare.F(true),
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

func TestDLPProfilePredefinedDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfilePredefinedDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestDLPProfilePredefinedGet(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Predefined.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfilePredefinedGetParams{
			AccountID: cloudflare.F("account_id"),
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

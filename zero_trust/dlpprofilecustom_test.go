// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
)

func TestDLPProfileCustomNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.New(context.TODO(), zero_trust.DLPProfileCustomNewParams{
		AccountID:           cloudflare.F("account_id"),
		Name:                cloudflare.F("name"),
		AIContextEnabled:    cloudflare.F(true),
		AllowedMatchCount:   cloudflare.F(int64(5)),
		ConfidenceThreshold: cloudflare.F("confidence_threshold"),
		ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
			Enabled: cloudflare.F(true),
			Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
				Files: cloudflare.F(true),
			}),
		}),
		Description: cloudflare.F("description"),
		Entries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsEntryUnion{zero_trust.DLPProfileCustomNewParamsEntriesDLPNewCustomEntry{
			Enabled: cloudflare.F(true),
			Name:    cloudflare.F("name"),
			Pattern: cloudflare.F(zero_trust.PatternParam{
				Regex:      cloudflare.F("regex"),
				Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
			}),
			Description: cloudflare.F("description"),
		}}),
		OCREnabled: cloudflare.F(true),
		SharedEntries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsSharedEntry{{
			Enabled: cloudflare.F(true),
			EntryID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestDLPProfileCustomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomUpdateParams{
			AccountID:           cloudflare.F("account_id"),
			Name:                cloudflare.F("name"),
			AIContextEnabled:    cloudflare.F(true),
			AllowedMatchCount:   cloudflare.F(int64(0)),
			ConfidenceThreshold: cloudflare.F("confidence_threshold"),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("description"),
			Entries: cloudflare.F([]zero_trust.DLPProfileCustomUpdateParamsEntryUnion{zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID{
				Enabled: cloudflare.F(true),
				EntryID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:    cloudflare.F("name"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("regex"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
				Description: cloudflare.F("description"),
			}}),
			OCREnabled: cloudflare.F(true),
			SharedEntries: cloudflare.F([]zero_trust.DLPProfileCustomUpdateParamsSharedEntry{{
				Enabled: cloudflare.F(true),
				EntryID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestDLPProfileCustomDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomDeleteParams{
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

func TestDLPProfileCustomGet(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.DLPProfileCustomGetParams{
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

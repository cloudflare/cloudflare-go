// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
)

func TestDLPProfileCustomNew(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.New(context.TODO(), zero_trust.DLPProfileCustomNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Profiles: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsProfile{{
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:       cloudflare.F("Generic CVV Card Number"),
			OCREnabled: cloudflare.F(true),
		}, {
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:       cloudflare.F("Generic CVV Card Number"),
			OCREnabled: cloudflare.F(true),
		}, {
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
			}}),
			Name:       cloudflare.F("Generic CVV Card Number"),
			OCREnabled: cloudflare.F(true),
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Update(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		zero_trust.DLPProfileCustomUpdateParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]zero_trust.DLPProfileCustomUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(zero_trust.PatternParam{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}}),
			Name:       cloudflare.F("Generic CVV Card Number"),
			OCREnabled: cloudflare.F(true),
			SharedEntries: cloudflare.F([]zero_trust.DLPProfileCustomUpdateParamsSharedEntryUnion{zero_trust.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}, zero_trust.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}, zero_trust.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Delete(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		zero_trust.DLPProfileCustomDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
	_, err := client.ZeroTrust.DLP.Profiles.Custom.Get(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		zero_trust.DLPProfileCustomGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

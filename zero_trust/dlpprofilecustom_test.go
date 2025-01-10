// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
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
		AccountID: cloudflare.F("account_id"),
		Body: zero_trust.DLPProfileCustomNewParamsBodyProfiles{
			Profiles: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsBodyProfilesProfile{{
				Entries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion{zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("name"),
					Pattern: cloudflare.F(zero_trust.PatternParam{
						Regex:      cloudflare.F("regex"),
						Validation: cloudflare.F(zero_trust.PatternValidationLuhn),
					}),
				}}),
				Name:                cloudflare.F("name"),
				AllowedMatchCount:   cloudflare.F(int64(5)),
				ConfidenceThreshold: cloudflare.F("confidence_threshold"),
				ContextAwareness: cloudflare.F(zero_trust.ContextAwarenessParam{
					Enabled: cloudflare.F(true),
					Skip: cloudflare.F(zero_trust.SkipConfigurationParam{
						Files: cloudflare.F(true),
					}),
				}),
				Description: cloudflare.F("description"),
				OCREnabled:  cloudflare.F(true),
				SharedEntries: cloudflare.F([]zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion{zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject{
					Enabled:   cloudflare.F(true),
					EntryID:   cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					EntryType: cloudflare.F(zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryTypeCustom),
				}}),
			}}),
		},
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
			}}),
			OCREnabled: cloudflare.F(true),
			SharedEntries: cloudflare.F([]zero_trust.DLPProfileCustomUpdateParamsSharedEntryUnion{zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObject{
				Enabled:   cloudflare.F(true),
				EntryID:   cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: cloudflare.F(zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined),
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

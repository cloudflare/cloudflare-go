// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZeroTrustDLPProfileCustomNew(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Customs.New(context.TODO(), cloudflare.ZeroTrustDLPProfileCustomNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Profiles: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomNewParamsProfile{{
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
		}, {
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
		}, {
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
				}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
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

func TestZeroTrustDLPProfileCustomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Customs.Update(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.ZeroTrustDLPProfileCustomUpdateParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
			SharedEntries: cloudflare.F([]cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntry{cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			})}),
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

func TestZeroTrustDLPProfileCustomDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Customs.Delete(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.ZeroTrustDLPProfileCustomDeleteParams{
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

func TestZeroTrustDLPProfileCustomGet(t *testing.T) {
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
	_, err := client.ZeroTrust.DLP.Profiles.Customs.Get(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.ZeroTrustDLPProfileCustomGetParams{
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

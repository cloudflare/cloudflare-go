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

func TestAccountDlpProfileCustomGet(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Dlp.Profiles.Customs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDlpProfileCustomUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Dlp.Profiles.Customs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountDlpProfileCustomUpdateParams{
			AllowedMatchCount: cloudflare.F(5.000000),
			Description:       cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.AccountDlpProfileCustomUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
			SharedEntries: cloudflare.F([]cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntry{cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined(cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined(cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined(cloudflare.AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined{
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

func TestAccountDlpProfileCustomDelete(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Dlp.Profiles.Customs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDlpProfileCustomDlpProfilesNewCustomProfiles(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Dlp.Profiles.Customs.DlpProfilesNewCustomProfiles(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParams{
			Profiles: cloudflare.F([]cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfile{{
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}}),
				Name: cloudflare.F("Generic CVV Card Number"),
			}, {
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}}),
				Name: cloudflare.F("Generic CVV Card Number"),
			}, {
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn),
					}),
				}}),
				Name: cloudflare.F("Generic CVV Card Number"),
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

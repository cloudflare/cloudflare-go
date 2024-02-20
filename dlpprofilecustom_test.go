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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.DLP.Profiles.Customs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DLPProfileCustomNewParams{
			Profiles: cloudflare.F([]cloudflare.DLPProfileCustomNewParamsProfile{{
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.DLPProfileCustomNewParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}}),
				Name: cloudflare.F("Generic CVV Card Number"),
			}, {
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.DLPProfileCustomNewParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}}),
				Name: cloudflare.F("Generic CVV Card Number"),
			}, {
				AllowedMatchCount: cloudflare.F(5.000000),
				Description:       cloudflare.F("A standard CVV card number"),
				Entries: cloudflare.F([]cloudflare.DLPProfileCustomNewParamsProfilesEntry{{
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
					}),
				}, {
					Enabled: cloudflare.F(true),
					Name:    cloudflare.F("Credit card (Visa)"),
					Pattern: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPattern{
						Regex:      cloudflare.F("^4[0-9]{6,14}$"),
						Validation: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn),
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.DLP.Profiles.Customs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.DLPProfileCustomUpdateParams{
			AllowedMatchCount: cloudflare.F(5.000000),
			Description:       cloudflare.F("A standard CVV card number"),
			Entries: cloudflare.F([]cloudflare.DLPProfileCustomUpdateParamsEntry{{
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}, {
				Enabled: cloudflare.F(true),
				Name:    cloudflare.F("Credit card (Visa)"),
				Pattern: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPattern{
					Regex:      cloudflare.F("^4[0-9]{6,14}$"),
					Validation: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsEntriesPatternValidationLuhn),
				}),
				ProfileID: cloudflare.F[any](map[string]interface{}{}),
			}}),
			Name: cloudflare.F("Generic CVV Card Number"),
			SharedEntries: cloudflare.F([]cloudflare.DLPProfileCustomUpdateParamsSharedEntry{cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
				Enabled: cloudflare.F(true),
			}), cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined(cloudflare.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined{
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.DLP.Profiles.Customs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.DLP.Profiles.Customs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

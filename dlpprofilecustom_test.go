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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DLP.Profiles.Customs.New(context.TODO(), cloudflare.DLPProfileCustomNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Profiles: cloudflare.F([]cloudflare.DLPProfileCustomNewParamsProfile{{
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
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
			ContextAwareness: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
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
			ContextAwareness: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.DLPProfileCustomNewParamsProfilesContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DLP.Profiles.Customs.Update(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.DLPProfileCustomUpdateParams{
			AccountID:         cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedMatchCount: cloudflare.F(5.000000),
			ContextAwareness: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsContextAwareness{
				Enabled: cloudflare.F(true),
				Skip: cloudflare.F(cloudflare.DLPProfileCustomUpdateParamsContextAwarenessSkip{
					Files: cloudflare.F(true),
				}),
			}),
			Description: cloudflare.F("A standard CVV card number"),
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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DLP.Profiles.Customs.Delete(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.DLPProfileCustomDeleteParams{
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.DLP.Profiles.Customs.Get(
		context.TODO(),
		"384e129d-25bd-403c-8019-bc19eb7a8a5f",
		cloudflare.DLPProfileCustomGetParams{
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

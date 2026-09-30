// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package basin_catalog_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/basin_catalog"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestMaintenanceConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.BasinCatalog.MaintenanceConfigs.Update(
		context.TODO(),
		"my-data-bucket",
		basin_catalog.MaintenanceConfigUpdateParams{
			AccountID: cloudflare.F("0123456789abcdef0123456789abcdef"),
			Compaction: cloudflare.F(basin_catalog.MaintenanceConfigUpdateParamsCompaction{
				State:        cloudflare.F(basin_catalog.MaintenanceConfigUpdateParamsCompactionStateEnabled),
				TargetSizeMB: cloudflare.F(basin_catalog.MaintenanceConfigUpdateParamsCompactionTargetSizeMB256),
			}),
			SnapshotExpiration: cloudflare.F(basin_catalog.MaintenanceConfigUpdateParamsSnapshotExpiration{
				MaxSnapshotAge:     cloudflare.F("14d"),
				MinSnapshotsToKeep: cloudflare.F(int64(5)),
				State:              cloudflare.F(basin_catalog.MaintenanceConfigUpdateParamsSnapshotExpirationStateEnabled),
			}),
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

func TestMaintenanceConfigGet(t *testing.T) {
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
	_, err := client.BasinCatalog.MaintenanceConfigs.Get(
		context.TODO(),
		"my-data-bucket",
		basin_catalog.MaintenanceConfigGetParams{
			AccountID: cloudflare.F("0123456789abcdef0123456789abcdef"),
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

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

func TestNamespaceTableMaintenanceConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.BasinCatalog.Namespaces.Tables.MaintenanceConfigs.Update(
		context.TODO(),
		"my-data-bucket",
		"my_namespace%1Fsub_namespace",
		"my_table",
		basin_catalog.NamespaceTableMaintenanceConfigUpdateParams{
			AccountID: cloudflare.F("0123456789abcdef0123456789abcdef"),
			Compaction: cloudflare.F(basin_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompaction{
				State:        cloudflare.F(basin_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompactionStateEnabled),
				TargetSizeMB: cloudflare.F(basin_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB256),
			}),
			SnapshotExpiration: cloudflare.F(basin_catalog.NamespaceTableMaintenanceConfigUpdateParamsSnapshotExpiration{
				MaxSnapshotAge:     cloudflare.F("14d"),
				MinSnapshotsToKeep: cloudflare.F(int64(5)),
				State:              cloudflare.F(basin_catalog.NamespaceTableMaintenanceConfigUpdateParamsSnapshotExpirationStateEnabled),
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

func TestNamespaceTableMaintenanceConfigGet(t *testing.T) {
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
	_, err := client.BasinCatalog.Namespaces.Tables.MaintenanceConfigs.Get(
		context.TODO(),
		"my-data-bucket",
		"my_namespace%1Fsub_namespace",
		"my_table",
		basin_catalog.NamespaceTableMaintenanceConfigGetParams{
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

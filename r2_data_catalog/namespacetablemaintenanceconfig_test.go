// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/r2_data_catalog"
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.R2DataCatalog.Namespaces.Tables.MaintenanceConfigs.Update(
		context.TODO(),
		"my-data-bucket",
		"my_namespace%1Fsub_namespace",
		"my_table",
		r2_data_catalog.NamespaceTableMaintenanceConfigUpdateParams{
			AccountID: cloudflare.F("0123456789abcdef0123456789abcdef"),
			Compaction: cloudflare.F(r2_data_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompaction{
				State:        cloudflare.F(r2_data_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompactionStateEnabled),
				TargetSizeMB: cloudflare.F(r2_data_catalog.NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB256),
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.R2DataCatalog.Namespaces.Tables.MaintenanceConfigs.Get(
		context.TODO(),
		"my-data-bucket",
		"my_namespace%1Fsub_namespace",
		"my_table",
		r2_data_catalog.NamespaceTableMaintenanceConfigGetParams{
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

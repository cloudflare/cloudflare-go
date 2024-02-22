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

func TestStorageKVNamespaceBulkUpdate(t *testing.T) {
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
	_, err := client.Storage.KV.Namespaces.Bulk.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
		cloudflare.StorageKVNamespaceBulkUpdateParams{
			Body: cloudflare.F([]cloudflare.StorageKVNamespaceBulkUpdateParamsBody{{
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F[any](map[string]interface{}{
					"someMetadataKey": "someMetadataValue",
				}),
				Value: cloudflare.F("Some string"),
			}, {
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F[any](map[string]interface{}{
					"someMetadataKey": "someMetadataValue",
				}),
				Value: cloudflare.F("Some string"),
			}, {
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F[any](map[string]interface{}{
					"someMetadataKey": "someMetadataValue",
				}),
				Value: cloudflare.F("Some string"),
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

func TestStorageKVNamespaceBulkDelete(t *testing.T) {
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
	_, err := client.Storage.KV.Namespaces.Bulk.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
		cloudflare.StorageKVNamespaceBulkDeleteParams{
			Body: cloudflare.F([]string{"My-Key", "My-Key", "My-Key"}),
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

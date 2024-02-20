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

func TestStorageKvNamespaceBulkDelete(t *testing.T) {
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
	_, err := client.Storage.Kv.Namespaces.Bulk.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
		cloudflare.StorageKvNamespaceBulkDeleteParams{
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

func TestStorageKvNamespaceBulkReplace(t *testing.T) {
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
	_, err := client.Storage.Kv.Namespaces.Bulk.Replace(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
		cloudflare.StorageKvNamespaceBulkReplaceParams{
			Body: cloudflare.F([]cloudflare.StorageKvNamespaceBulkReplaceParamsBody{{
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

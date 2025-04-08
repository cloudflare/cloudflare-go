// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/secrets_store"
)

func TestStoreSecretNew(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		secrets_store.StoreSecretNewParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
			Body: []secrets_store.StoreSecretNewParamsBody{{
				Name:  cloudflare.F("MY_API_KEY"),
				Value: cloudflare.F("api-token-secret-123"),
			}},
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

func TestStoreSecretListWithOptionalParams(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		secrets_store.StoreSecretListParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
			Direction: cloudflare.F(secrets_store.StoreSecretListParamsDirectionAsc),
			Order:     cloudflare.F(secrets_store.StoreSecretListParamsOrderName),
			Page:      cloudflare.F(int64(2)),
			PerPage:   cloudflare.F(int64(20)),
			Search:    cloudflare.F("search"),
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

func TestStoreSecretDelete(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
		secrets_store.StoreSecretDeleteParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
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

func TestStoreSecretBulkDelete(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.BulkDelete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		secrets_store.StoreSecretBulkDeleteParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
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

func TestStoreSecretBulkdEdit(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.BulkdEdit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		secrets_store.StoreSecretBulkdEditParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
			Body: []secrets_store.StoreSecretBulkdEditParamsBody{{
				Name:  cloudflare.F("MY_API_KEY"),
				Value: cloudflare.F("api-token-secret-123"),
			}},
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

func TestStoreSecretDuplicate(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.Duplicate(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
		secrets_store.StoreSecretDuplicateParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("MY_API_KEY"),
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

func TestStoreSecretEditWithOptionalParams(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
		secrets_store.StoreSecretEditParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("MY_API_KEY"),
			Value:     cloudflare.F("api-token-secret-123"),
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

func TestStoreSecretGet(t *testing.T) {
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
	_, err := client.SecretsStore.Stores.Secrets.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
		secrets_store.StoreSecretGetParams{
			AccountID: cloudflare.F("985e105f4ecef8ad9ca31a8372d0c353"),
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
)

func TestConnectivityDirectoryServiceNewWithOptionalParams(t *testing.T) {
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
	err := client.ZeroTrust.Connectivity.Directory.Services.New(context.TODO(), zero_trust.ConnectivityDirectoryServiceNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Host: cloudflare.F(zero_trust.ConnectivityDirectoryServiceNewParamsHost{
			Hostname:        cloudflare.F("hostname"),
			IPV4:            cloudflare.F("ipv4"),
			IPV6:            cloudflare.F("ipv6"),
			Network:         cloudflare.F[any](map[string]interface{}{}),
			ResolverNetwork: cloudflare.F[any](map[string]interface{}{}),
		}),
		Name:      cloudflare.F("name"),
		Type:      cloudflare.F(zero_trust.ConnectivityDirectoryServiceNewParamsTypeHTTP),
		HTTPPort:  cloudflare.F(int64(1)),
		HTTPSPort: cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectivityDirectoryServiceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.ZeroTrust.Connectivity.Directory.Services.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.ConnectivityDirectoryServiceUpdateParams{
			AccountID: cloudflare.F("account_id"),
			Host: cloudflare.F(zero_trust.ConnectivityDirectoryServiceUpdateParamsHost{
				Hostname:        cloudflare.F("hostname"),
				IPV4:            cloudflare.F("ipv4"),
				IPV6:            cloudflare.F("ipv6"),
				Network:         cloudflare.F[any](map[string]interface{}{}),
				ResolverNetwork: cloudflare.F[any](map[string]interface{}{}),
			}),
			Name:      cloudflare.F("name"),
			Type:      cloudflare.F(zero_trust.ConnectivityDirectoryServiceUpdateParamsTypeHTTP),
			HTTPPort:  cloudflare.F(int64(1)),
			HTTPSPort: cloudflare.F(int64(1)),
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

func TestConnectivityDirectoryServiceListWithOptionalParams(t *testing.T) {
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
	err := client.ZeroTrust.Connectivity.Directory.Services.List(context.TODO(), zero_trust.ConnectivityDirectoryServiceListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
		Type:      cloudflare.F(zero_trust.ConnectivityDirectoryServiceListParamsTypeHTTP),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectivityDirectoryServiceDelete(t *testing.T) {
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
	err := client.ZeroTrust.Connectivity.Directory.Services.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.ConnectivityDirectoryServiceDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestConnectivityDirectoryServiceGet(t *testing.T) {
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
	err := client.ZeroTrust.Connectivity.Directory.Services.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.ConnectivityDirectoryServiceGetParams{
			AccountID: cloudflare.F("account_id"),
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

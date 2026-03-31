// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/connectivity"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

func TestDirectoryServiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectivity.Directory.Services.New(context.TODO(), connectivity.DirectoryServiceNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfig{
			Host: cloudflare.F[connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion](connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost{
				Hostname: cloudflare.F("api.example.com"),
				ResolverNetwork: cloudflare.F(connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork{
					TunnelID:    cloudflare.F("0191dce4-9ab4-7fce-b660-8e5dec5172da"),
					ResolverIPs: cloudflare.F([]string{"string"}),
				}),
			}),
			Name:      cloudflare.F("web-server"),
			Type:      cloudflare.F(connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTypeHTTP),
			HTTPPort:  cloudflare.F(int64(8080)),
			HTTPSPort: cloudflare.F(int64(8443)),
			TLSSettings: cloudflare.F(connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTLSSettings{
				CERTVerificationMode: cloudflare.F("verify_full"),
			}),
		},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirectoryServiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectivity.Directory.Services.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		connectivity.DirectoryServiceUpdateParams{
			AccountID: cloudflare.F("account_id"),
			Body: connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfig{
				Host: cloudflare.F[connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion](connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host{
					IPV4: cloudflare.F("10.0.0.1"),
					Network: cloudflare.F(connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork{
						TunnelID: cloudflare.F("0191dce4-9ab4-7fce-b660-8e5dec5172da"),
					}),
				}),
				Name:      cloudflare.F("postgres-db"),
				Type:      cloudflare.F(connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTypeTCP),
				HTTPPort:  cloudflare.F(int64(8080)),
				HTTPSPort: cloudflare.F(int64(8443)),
				TLSSettings: cloudflare.F(connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTLSSettings{
					CERTVerificationMode: cloudflare.F("verify_full"),
				}),
			},
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

func TestDirectoryServiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectivity.Directory.Services.List(context.TODO(), connectivity.DirectoryServiceListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
		Type:      cloudflare.F(connectivity.DirectoryServiceListParamsTypeTCP),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirectoryServiceDelete(t *testing.T) {
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
	err := client.Connectivity.Directory.Services.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		connectivity.DirectoryServiceDeleteParams{
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

func TestDirectoryServiceGet(t *testing.T) {
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
	_, err := client.Connectivity.Directory.Services.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		connectivity.DirectoryServiceGetParams{
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

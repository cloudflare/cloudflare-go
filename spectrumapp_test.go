// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestSpectrumAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrums.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		cloudflare.SpectrumAppUpdateParams{
			DNS: cloudflare.F(cloudflare.SpectrumAppUpdateParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(cloudflare.SpectrumAppUpdateParamsDNSTypeCname),
			}),
			OriginDNS: cloudflare.F(cloudflare.SpectrumAppUpdateParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(cloudflare.SpectrumAppUpdateParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[cloudflare.SpectrumAppUpdateParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[cloudflare.SpectrumAppUpdateParamsEdgeIPs](cloudflare.SpectrumAppUpdateParamsEdgeIPsObject(cloudflare.SpectrumAppUpdateParamsEdgeIPsObject{
				Connectivity: cloudflare.F(cloudflare.SpectrumAppUpdateParamsEdgeIPsObjectConnectivityAll),
				Type:         cloudflare.F(cloudflare.SpectrumAppUpdateParamsEdgeIPsObjectTypeDynamic),
			})),
			IPFirewall:    cloudflare.F(true),
			ProxyProtocol: cloudflare.F(cloudflare.SpectrumAppUpdateParamsProxyProtocolOff),
			TLS:           cloudflare.F(cloudflare.SpectrumAppUpdateParamsTLSFull),
			TrafficType:   cloudflare.F(cloudflare.SpectrumAppUpdateParamsTrafficTypeDirect),
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

func TestSpectrumAppDelete(t *testing.T) {
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
	_, err := client.Spectrums.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpectrumAppGet(t *testing.T) {
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
	_, err := client.Spectrums.Apps.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrums.Apps.SpectrumApplicationsNewSpectrumApplicationUsingANameForTheOrigin(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParams{
			DNS: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSTypeCname),
			}),
			OriginDNS: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPs](cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject{
				Connectivity: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivityAll),
				Type:         cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectTypeDynamic),
			})),
			IPFirewall:    cloudflare.F(true),
			ProxyProtocol: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocolOff),
			TLS:           cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLSFull),
			TrafficType:   cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficTypeDirect),
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

func TestSpectrumAppSpectrumApplicationsListSpectrumApplicationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrums.Apps.SpectrumApplicationsListSpectrumApplications(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.SpectrumAppSpectrumApplicationsListSpectrumApplicationsParams{
			Direction: cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirectionDesc),
			Order:     cloudflare.F(cloudflare.SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderProtocol),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(1.000000),
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

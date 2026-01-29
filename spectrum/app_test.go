// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/cloudflare/cloudflare-go/v6/spectrum"
)

func TestAppNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate auth errors on test suite")
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
	_, err := client.Spectrum.Apps.New(context.TODO(), spectrum.AppNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: spectrum.AppNewParamsBodySpectrumConfigAppConfig{
			DNS: cloudflare.F(spectrum.DNSParam{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(spectrum.DNSTypeCNAME),
			}),
			Protocol:         cloudflare.F("tcp/22"),
			TrafficType:      cloudflare.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigTrafficTypeDirect),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[spectrum.EdgeIPsUnionParam](spectrum.EdgeIPsDynamicParam{
				Connectivity: cloudflare.F(spectrum.EdgeIPsDynamicConnectivityAll),
				Type:         cloudflare.F(spectrum.EdgeIPsDynamicTypeDynamic),
			}),
			IPFirewall:   cloudflare.F(false),
			OriginDirect: cloudflare.F([]string{"tcp://127.0.0.1:8080"}),
			OriginDNS: cloudflare.F(spectrum.OriginDNSParam{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(spectrum.OriginDNSTypeEmpty),
			}),
			OriginPort:    cloudflare.F[spectrum.OriginPortUnionParam](shared.UnionInt(int64(22))),
			ProxyProtocol: cloudflare.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigProxyProtocolOff),
			TLS:           cloudflare.F(spectrum.AppNewParamsBodySpectrumConfigAppConfigTLSOff),
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

func TestAppUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate auth errors on test suite")
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
	_, err := client.Spectrum.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: spectrum.AppUpdateParamsBodySpectrumConfigAppConfig{
				DNS: cloudflare.F(spectrum.DNSParam{
					Name: cloudflare.F("ssh.example.com"),
					Type: cloudflare.F(spectrum.DNSTypeCNAME),
				}),
				Protocol:         cloudflare.F("tcp/22"),
				TrafficType:      cloudflare.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeDirect),
				ArgoSmartRouting: cloudflare.F(true),
				EdgeIPs: cloudflare.F[spectrum.EdgeIPsUnionParam](spectrum.EdgeIPsDynamicParam{
					Connectivity: cloudflare.F(spectrum.EdgeIPsDynamicConnectivityAll),
					Type:         cloudflare.F(spectrum.EdgeIPsDynamicTypeDynamic),
				}),
				IPFirewall:   cloudflare.F(false),
				OriginDirect: cloudflare.F([]string{"tcp://127.0.0.1:8080"}),
				OriginDNS: cloudflare.F(spectrum.OriginDNSParam{
					Name: cloudflare.F("origin.example.com"),
					TTL:  cloudflare.F(int64(600)),
					Type: cloudflare.F(spectrum.OriginDNSTypeEmpty),
				}),
				OriginPort:    cloudflare.F[spectrum.OriginPortUnionParam](shared.UnionInt(int64(22))),
				ProxyProtocol: cloudflare.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolOff),
				TLS:           cloudflare.F(spectrum.AppUpdateParamsBodySpectrumConfigAppConfigTLSOff),
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

func TestAppListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate HTTP 422 errors on test suite")
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
	_, err := client.Spectrum.Apps.List(context.TODO(), spectrum.AppListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(spectrum.AppListParamsDirectionDesc),
		Order:     cloudflare.F(spectrum.AppListParamsOrderProtocol),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(1.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppDelete(t *testing.T) {
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
	_, err := client.Spectrum.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestAppGet(t *testing.T) {
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
	_, err := client.Spectrum.Apps.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

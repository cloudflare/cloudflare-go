// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/spectrum"
)

func TestAppNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Spectrum.Apps.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppNewParams{
			DNS: cloudflare.F(spectrum.AppNewParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(spectrum.AppNewParamsDNSTypeCNAME),
			}),
			OriginDNS: cloudflare.F(spectrum.AppNewParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(spectrum.AppNewParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[spectrum.AppNewParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[spectrum.AppNewParamsEdgeIPs](spectrum.AppNewParamsEdgeIPsEyeballIPs(spectrum.AppNewParamsEdgeIPsEyeballIPs{
				Connectivity: cloudflare.F(spectrum.AppNewParamsEdgeIPsEyeballIPsConnectivityAll),
				Type:         cloudflare.F(spectrum.AppNewParamsEdgeIPsEyeballIPsTypeDynamic),
			})),
			IPFirewall:    cloudflare.F(true),
			ProxyProtocol: cloudflare.F(spectrum.AppNewParamsProxyProtocolOff),
			TLS:           cloudflare.F(spectrum.AppNewParamsTLSFull),
			TrafficType:   cloudflare.F(spectrum.AppNewParamsTrafficTypeDirect),
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

func TestAppUpdateWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Spectrum.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		spectrum.AppUpdateParams{
			DNS: cloudflare.F(spectrum.AppUpdateParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(spectrum.AppUpdateParamsDNSTypeCNAME),
			}),
			OriginDNS: cloudflare.F(spectrum.AppUpdateParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(spectrum.AppUpdateParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[spectrum.AppUpdateParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[spectrum.AppUpdateParamsEdgeIPs](spectrum.AppUpdateParamsEdgeIPsEyeballIPs(spectrum.AppUpdateParamsEdgeIPsEyeballIPs{
				Connectivity: cloudflare.F(spectrum.AppUpdateParamsEdgeIPsEyeballIPsConnectivityAll),
				Type:         cloudflare.F(spectrum.AppUpdateParamsEdgeIPsEyeballIPsTypeDynamic),
			})),
			IPFirewall:    cloudflare.F(true),
			ProxyProtocol: cloudflare.F(spectrum.AppUpdateParamsProxyProtocolOff),
			TLS:           cloudflare.F(spectrum.AppUpdateParamsTLSFull),
			TrafficType:   cloudflare.F(spectrum.AppUpdateParamsTrafficTypeDirect),
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
	)
	_, err := client.Spectrum.Apps.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AppListParams{
			Direction: cloudflare.F(spectrum.AppListParamsDirectionDesc),
			Order:     cloudflare.F(spectrum.AppListParamsOrderProtocol),
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

func TestAppDelete(t *testing.T) {
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
	)
	_, err := client.Spectrum.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		spectrum.AppDeleteParams{
			Body: cloudflare.F[any](map[string]interface{}{}),
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
	)
	_, err := client.Spectrum.Apps.Get(
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

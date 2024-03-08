// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestSpectrumAppNewWithOptionalParams(t *testing.T) {
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
		cloudflare.SpectrumAppNewParams{
			DNS: cloudflare.F(cloudflare.SpectrumAppNewParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(cloudflare.SpectrumAppNewParamsDNSTypeCNAME),
			}),
			OriginDNS: cloudflare.F(cloudflare.SpectrumAppNewParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(cloudflare.SpectrumAppNewParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[cloudflare.SpectrumAppNewParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[cloudflare.SpectrumAppNewParamsEdgeIPs](cloudflare.SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable(cloudflare.SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable{
				Connectivity: cloudflare.F(cloudflare.SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll),
				Type:         cloudflare.F(cloudflare.SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic),
			})),
			IPFirewall:    cloudflare.F(true),
			ProxyProtocol: cloudflare.F(cloudflare.SpectrumAppNewParamsProxyProtocolOff),
			TLS:           cloudflare.F(cloudflare.SpectrumAppNewParamsTLSFull),
			TrafficType:   cloudflare.F(cloudflare.SpectrumAppNewParamsTrafficTypeDirect),
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Spectrum.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		cloudflare.SpectrumAppUpdateParams{
			DNS: cloudflare.F(cloudflare.SpectrumAppUpdateParamsDNS{
				Name: cloudflare.F("ssh.example.com"),
				Type: cloudflare.F(cloudflare.SpectrumAppUpdateParamsDNSTypeCNAME),
			}),
			OriginDNS: cloudflare.F(cloudflare.SpectrumAppUpdateParamsOriginDNS{
				Name: cloudflare.F("origin.example.com"),
				TTL:  cloudflare.F(int64(600)),
				Type: cloudflare.F(cloudflare.SpectrumAppUpdateParamsOriginDNSTypeEmpty),
			}),
			OriginPort:       cloudflare.F[cloudflare.SpectrumAppUpdateParamsOriginPort](shared.UnionInt(int64(22))),
			Protocol:         cloudflare.F("tcp/22"),
			ArgoSmartRouting: cloudflare.F(true),
			EdgeIPs: cloudflare.F[cloudflare.SpectrumAppUpdateParamsEdgeIPs](cloudflare.SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable(cloudflare.SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable{
				Connectivity: cloudflare.F(cloudflare.SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll),
				Type:         cloudflare.F(cloudflare.SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic),
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

func TestSpectrumAppListWithOptionalParams(t *testing.T) {
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
		cloudflare.SpectrumAppListParams{
			Direction: cloudflare.F(cloudflare.SpectrumAppListParamsDirectionDesc),
			Order:     cloudflare.F(cloudflare.SpectrumAppListParamsOrderProtocol),
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Spectrum.Apps.Delete(
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

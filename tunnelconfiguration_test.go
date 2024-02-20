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

func TestTunnelConfigurationList(t *testing.T) {
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
	_, err := client.Tunnels.Configurations.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTunnelConfigurationReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Tunnels.Configurations.Replace(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		cloudflare.TunnelConfigurationReplaceParams{
			Config: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfig{
				Ingress: cloudflare.F([]cloudflare.TunnelConfigurationReplaceParamsConfigIngress{{
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}}),
				OriginRequest: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigOriginRequest{
					Access: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigOriginRequestAccess{
						AudTag:   cloudflare.F([]string{"string", "string", "string"}),
						Required: cloudflare.F(true),
						TeamName: cloudflare.F("string"),
					}),
					CaPool:                 cloudflare.F("string"),
					ConnectTimeout:         cloudflare.F(int64(0)),
					DisableChunkedEncoding: cloudflare.F(true),
					HTTP2Origin:            cloudflare.F(true),
					HTTPHostHeader:         cloudflare.F("string"),
					KeepAliveConnections:   cloudflare.F(int64(0)),
					KeepAliveTimeout:       cloudflare.F(int64(0)),
					NoHappyEyeballs:        cloudflare.F(true),
					NoTLSVerify:            cloudflare.F(true),
					OriginServerName:       cloudflare.F("string"),
					ProxyType:              cloudflare.F("string"),
					TcpKeepAlive:           cloudflare.F(int64(0)),
					TLSTimeout:             cloudflare.F(int64(0)),
				}),
				WarpRouting: cloudflare.F(cloudflare.TunnelConfigurationReplaceParamsConfigWarpRouting{
					Enabled: cloudflare.F(true),
				}),
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

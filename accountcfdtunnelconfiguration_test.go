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

func TestAccountCfdTunnelConfigurationCloudflareTunnelConfigurationGetConfiguration(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.CfdTunnels.Configurations.CloudflareTunnelConfigurationGetConfiguration(
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

func TestAccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.CfdTunnels.Configurations.CloudflareTunnelConfigurationPutConfiguration(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams{
			Config: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfig{
				Ingress: cloudflare.F([]cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngress{{
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						Http2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTlsVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TlsTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						Http2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTlsVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TlsTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					OriginRequest: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequest{
						Access: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigIngressOriginRequestAccess{
							AudTag:   cloudflare.F([]string{"string", "string", "string"}),
							Required: cloudflare.F(true),
							TeamName: cloudflare.F("string"),
						}),
						CaPool:                 cloudflare.F("string"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						Http2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("string"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTlsVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("string"),
						ProxyType:              cloudflare.F("string"),
						TcpKeepAlive:           cloudflare.F(int64(0)),
						TlsTimeout:             cloudflare.F(int64(0)),
					}),
					Path:    cloudflare.F("subpath"),
					Service: cloudflare.F("https://localhost:8001"),
				}}),
				OriginRequest: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequest{
					Access: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigOriginRequestAccess{
						AudTag:   cloudflare.F([]string{"string", "string", "string"}),
						Required: cloudflare.F(true),
						TeamName: cloudflare.F("string"),
					}),
					CaPool:                 cloudflare.F("string"),
					ConnectTimeout:         cloudflare.F(int64(0)),
					DisableChunkedEncoding: cloudflare.F(true),
					Http2Origin:            cloudflare.F(true),
					HTTPHostHeader:         cloudflare.F("string"),
					KeepAliveConnections:   cloudflare.F(int64(0)),
					KeepAliveTimeout:       cloudflare.F(int64(0)),
					NoHappyEyeballs:        cloudflare.F(true),
					NoTlsVerify:            cloudflare.F(true),
					OriginServerName:       cloudflare.F("string"),
					ProxyType:              cloudflare.F("string"),
					TcpKeepAlive:           cloudflare.F(int64(0)),
					TlsTimeout:             cloudflare.F(int64(0)),
				}),
				WarpRouting: cloudflare.F(cloudflare.AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParamsConfigWarpRouting{
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

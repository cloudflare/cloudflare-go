// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestTunnelCloudflaredConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Cloudflared.Configurations.Update(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelCloudflaredConfigurationUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Config: cloudflare.F(zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfig{
				Ingress: cloudflare.F([]zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfigIngress{{
					Hostname: cloudflare.F("tunnel.example.com"),
					Service:  cloudflare.F("https://localhost:8001"),
					OriginRequest: cloudflare.F(zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequest{
						Access: cloudflare.F(zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequestAccess{
							AUDTag:   cloudflare.F([]string{"string"}),
							TeamName: cloudflare.F("zero-trust-organization-name"),
							Required: cloudflare.F(false),
						}),
						CAPool:                 cloudflare.F("caPool"),
						ConnectTimeout:         cloudflare.F(int64(10)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("httpHostHeader"),
						KeepAliveConnections:   cloudflare.F(int64(100)),
						KeepAliveTimeout:       cloudflare.F(int64(90)),
						NoHappyEyeballs:        cloudflare.F(false),
						NoTLSVerify:            cloudflare.F(false),
						OriginServerName:       cloudflare.F("originServerName"),
						ProxyType:              cloudflare.F("proxyType"),
						TCPKeepAlive:           cloudflare.F(int64(30)),
						TLSTimeout:             cloudflare.F(int64(10)),
					}),
					Path: cloudflare.F("subpath"),
				}}),
				OriginRequest: cloudflare.F(zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequest{
					Access: cloudflare.F(zero_trust.TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequestAccess{
						AUDTag:   cloudflare.F([]string{"string"}),
						TeamName: cloudflare.F("zero-trust-organization-name"),
						Required: cloudflare.F(false),
					}),
					CAPool:                 cloudflare.F("caPool"),
					ConnectTimeout:         cloudflare.F(int64(10)),
					DisableChunkedEncoding: cloudflare.F(true),
					HTTP2Origin:            cloudflare.F(true),
					HTTPHostHeader:         cloudflare.F("httpHostHeader"),
					KeepAliveConnections:   cloudflare.F(int64(100)),
					KeepAliveTimeout:       cloudflare.F(int64(90)),
					NoHappyEyeballs:        cloudflare.F(false),
					NoTLSVerify:            cloudflare.F(false),
					OriginServerName:       cloudflare.F("originServerName"),
					ProxyType:              cloudflare.F("proxyType"),
					TCPKeepAlive:           cloudflare.F(int64(30)),
					TLSTimeout:             cloudflare.F(int64(10)),
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

func TestTunnelCloudflaredConfigurationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Tunnels.Cloudflared.Configurations.Get(
		context.TODO(),
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelCloudflaredConfigurationGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

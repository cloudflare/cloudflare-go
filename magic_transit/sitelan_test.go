// File generated from our OpenAPI spec by Stainless.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/magic_transit"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestSiteLanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Lans.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLanNewParams{
			Lan: cloudflare.F(magic_transit.SiteLanNewParamsLan{
				Description: cloudflare.F("string"),
				HaLink:      cloudflare.F(true),
				Nat: cloudflare.F(magic_transit.SiteLanNewParamsLanNat{
					StaticPrefix: cloudflare.F("192.0.2.0/24"),
				}),
				Physport: cloudflare.F(int64(1)),
				RoutedSubnets: cloudflare.F([]magic_transit.SiteLanNewParamsLanRoutedSubnet{{
					Nat: cloudflare.F(magic_transit.SiteLanNewParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}, {
					Nat: cloudflare.F(magic_transit.SiteLanNewParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}, {
					Nat: cloudflare.F(magic_transit.SiteLanNewParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}}),
				StaticAddressing: cloudflare.F(magic_transit.SiteLanNewParamsLanStaticAddressing{
					Address: cloudflare.F("192.0.2.0/24"),
					DhcpRelay: cloudflare.F(magic_transit.SiteLanNewParamsLanStaticAddressingDhcpRelay{
						ServerAddresses: cloudflare.F([]string{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
					}),
					DhcpServer: cloudflare.F(magic_transit.SiteLanNewParamsLanStaticAddressingDhcpServer{
						DhcpPoolEnd:   cloudflare.F("192.0.2.1"),
						DhcpPoolStart: cloudflare.F("192.0.2.1"),
						DNSServer:     cloudflare.F("192.0.2.1"),
						Reservations: cloudflare.F(map[string]string{
							"00:11:22:33:44:55": "192.0.2.100",
							"AA:BB:CC:DD:EE:FF": "192.168.1.101",
						}),
					}),
					SecondaryAddress: cloudflare.F("192.0.2.0/24"),
					VirtualAddress:   cloudflare.F("192.0.2.0/24"),
				}),
				VlanTag: cloudflare.F(int64(0)),
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

func TestSiteLanUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Lans.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLanUpdateParams{
			Lan: cloudflare.F(magic_transit.SiteLanUpdateParamsLan{
				Description: cloudflare.F("string"),
				Nat: cloudflare.F(magic_transit.SiteLanUpdateParamsLanNat{
					StaticPrefix: cloudflare.F("192.0.2.0/24"),
				}),
				Physport: cloudflare.F(int64(1)),
				RoutedSubnets: cloudflare.F([]magic_transit.SiteLanUpdateParamsLanRoutedSubnet{{
					Nat: cloudflare.F(magic_transit.SiteLanUpdateParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}, {
					Nat: cloudflare.F(magic_transit.SiteLanUpdateParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}, {
					Nat: cloudflare.F(magic_transit.SiteLanUpdateParamsLanRoutedSubnetsNat{
						StaticPrefix: cloudflare.F("192.0.2.0/24"),
					}),
					NextHop: cloudflare.F("192.0.2.1"),
					Prefix:  cloudflare.F("192.0.2.0/24"),
				}}),
				StaticAddressing: cloudflare.F(magic_transit.SiteLanUpdateParamsLanStaticAddressing{
					Address: cloudflare.F("192.0.2.0/24"),
					DhcpRelay: cloudflare.F(magic_transit.SiteLanUpdateParamsLanStaticAddressingDhcpRelay{
						ServerAddresses: cloudflare.F([]string{"192.0.2.1", "192.0.2.1", "192.0.2.1"}),
					}),
					DhcpServer: cloudflare.F(magic_transit.SiteLanUpdateParamsLanStaticAddressingDhcpServer{
						DhcpPoolEnd:   cloudflare.F("192.0.2.1"),
						DhcpPoolStart: cloudflare.F("192.0.2.1"),
						DNSServer:     cloudflare.F("192.0.2.1"),
						Reservations: cloudflare.F(map[string]string{
							"00:11:22:33:44:55": "192.0.2.100",
							"AA:BB:CC:DD:EE:FF": "192.168.1.101",
						}),
					}),
					SecondaryAddress: cloudflare.F("192.0.2.0/24"),
					VirtualAddress:   cloudflare.F("192.0.2.0/24"),
				}),
				VlanTag: cloudflare.F(int64(0)),
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

func TestSiteLanList(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Lans.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteLanDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Lans.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSiteLanGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.Lans.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

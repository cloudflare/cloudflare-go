// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestSiteLANNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Physport:  cloudflare.F(int64(1)),
			VlanTag:   cloudflare.F(int64(0)),
			HaLink:    cloudflare.F(true),
			Name:      cloudflare.F("name"),
			Nat: cloudflare.F(magic_transit.NatParam{
				StaticPrefix: cloudflare.F("192.0.2.0/24"),
			}),
			RoutedSubnets: cloudflare.F([]magic_transit.RoutedSubnetParam{{
				NextHop: cloudflare.F("192.0.2.1"),
				Prefix:  cloudflare.F("192.0.2.0/24"),
				Nat: cloudflare.F(magic_transit.NatParam{
					StaticPrefix: cloudflare.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: cloudflare.F(magic_transit.LANStaticAddressingParam{
				Address: cloudflare.F("192.0.2.0/24"),
				DHCPRelay: cloudflare.F(magic_transit.DHCPRelayParam{
					ServerAddresses: cloudflare.F([]string{"192.0.2.1"}),
				}),
				DHCPServer: cloudflare.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   cloudflare.F("192.0.2.1"),
					DHCPPoolStart: cloudflare.F("192.0.2.1"),
					DNSServer:     cloudflare.F("192.0.2.1"),
					Reservations: cloudflare.F(map[string]string{
						"00:11:22:33:44:55": "192.0.2.100",
						"AA:BB:CC:DD:EE:FF": "192.168.1.101",
					}),
				}),
				SecondaryAddress: cloudflare.F("192.0.2.0/24"),
				VirtualAddress:   cloudflare.F("192.0.2.0/24"),
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

func TestSiteLANUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("name"),
			Nat: cloudflare.F(magic_transit.NatParam{
				StaticPrefix: cloudflare.F("192.0.2.0/24"),
			}),
			Physport: cloudflare.F(int64(1)),
			RoutedSubnets: cloudflare.F([]magic_transit.RoutedSubnetParam{{
				NextHop: cloudflare.F("192.0.2.1"),
				Prefix:  cloudflare.F("192.0.2.0/24"),
				Nat: cloudflare.F(magic_transit.NatParam{
					StaticPrefix: cloudflare.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: cloudflare.F(magic_transit.LANStaticAddressingParam{
				Address: cloudflare.F("192.0.2.0/24"),
				DHCPRelay: cloudflare.F(magic_transit.DHCPRelayParam{
					ServerAddresses: cloudflare.F([]string{"192.0.2.1"}),
				}),
				DHCPServer: cloudflare.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   cloudflare.F("192.0.2.1"),
					DHCPPoolStart: cloudflare.F("192.0.2.1"),
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

func TestSiteLANList(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANListParams{
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

func TestSiteLANDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANDeleteParams{
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

func TestSiteLANEditWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANEditParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("name"),
			Nat: cloudflare.F(magic_transit.NatParam{
				StaticPrefix: cloudflare.F("192.0.2.0/24"),
			}),
			Physport: cloudflare.F(int64(1)),
			RoutedSubnets: cloudflare.F([]magic_transit.RoutedSubnetParam{{
				NextHop: cloudflare.F("192.0.2.1"),
				Prefix:  cloudflare.F("192.0.2.0/24"),
				Nat: cloudflare.F(magic_transit.NatParam{
					StaticPrefix: cloudflare.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: cloudflare.F(magic_transit.LANStaticAddressingParam{
				Address: cloudflare.F("192.0.2.0/24"),
				DHCPRelay: cloudflare.F(magic_transit.DHCPRelayParam{
					ServerAddresses: cloudflare.F([]string{"192.0.2.1"}),
				}),
				DHCPServer: cloudflare.F(magic_transit.DHCPServerParam{
					DHCPPoolEnd:   cloudflare.F("192.0.2.1"),
					DHCPPoolStart: cloudflare.F("192.0.2.1"),
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

func TestSiteLANGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.LANs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteLANGetParams{
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

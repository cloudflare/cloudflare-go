// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/magic_transit"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestSiteACLNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			LAN1: cloudflare.F(magic_transit.ACLConfigurationParam{
				LANID:   cloudflare.F("string"),
				LANName: cloudflare.F("string"),
				Ports:   cloudflare.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: cloudflare.F([]magic_transit.SubnetUnionParam{shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1")}),
			}),
			LAN2: cloudflare.F(magic_transit.ACLConfigurationParam{
				LANID:   cloudflare.F("string"),
				LANName: cloudflare.F("string"),
				Ports:   cloudflare.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: cloudflare.F([]magic_transit.SubnetUnionParam{shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1")}),
			}),
			Name:           cloudflare.F("PIN Pad - Cash Register"),
			Description:    cloudflare.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: cloudflare.F(true),
			Protocols:      cloudflare.F([]magic_transit.AllowedProtocol{magic_transit.AllowedProtocolTCP, magic_transit.AllowedProtocolUdp, magic_transit.AllowedProtocolIcmp}),
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

func TestSiteACLUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLUpdateParams{
			AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Description:    cloudflare.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: cloudflare.F(true),
			LAN1: cloudflare.F(magic_transit.ACLConfigurationParam{
				LANID:   cloudflare.F("string"),
				LANName: cloudflare.F("string"),
				Ports:   cloudflare.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: cloudflare.F([]magic_transit.SubnetUnionParam{shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1")}),
			}),
			LAN2: cloudflare.F(magic_transit.ACLConfigurationParam{
				LANID:   cloudflare.F("string"),
				LANName: cloudflare.F("string"),
				Ports:   cloudflare.F([]int64{int64(1), int64(1), int64(1)}),
				Subnets: cloudflare.F([]magic_transit.SubnetUnionParam{shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1"), shared.UnionString("192.0.2.1")}),
			}),
			Name:      cloudflare.F("PIN Pad - Cash Register"),
			Protocols: cloudflare.F([]magic_transit.AllowedProtocol{magic_transit.AllowedProtocolTCP, magic_transit.AllowedProtocolUdp, magic_transit.AllowedProtocolIcmp}),
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

func TestSiteACLList(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLListParams{
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

func TestSiteACLDelete(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLDeleteParams{
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

func TestSiteACLGet(t *testing.T) {
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
	_, err := client.MagicTransit.Sites.ACLs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		magic_transit.SiteACLGetParams{
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

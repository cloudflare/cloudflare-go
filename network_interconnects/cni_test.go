// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/network_interconnects"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestCNINewWithOptionalParams(t *testing.T) {
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
	_, err := client.NetworkInterconnects.CNIs.New(context.TODO(), network_interconnects.CNINewParams{
		AccountID:    cloudflare.F("account_id"),
		Account:      cloudflare.F("account"),
		Interconnect: cloudflare.F("interconnect"),
		Magic: cloudflare.F(network_interconnects.CNINewParamsMagic{
			ConduitName: cloudflare.F("conduit_name"),
			Description: cloudflare.F("description"),
			Mtu:         cloudflare.F(int64(0)),
		}),
		BGP: cloudflare.F(network_interconnects.CNINewParamsBGP{
			CustomerASN:   cloudflare.F(int64(0)),
			ExtraPrefixes: cloudflare.F([]string{"string"}),
			Md5Key:        cloudflare.F("md5_key"),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCNIUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.NetworkInterconnects.CNIs.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		network_interconnects.CNIUpdateParams{
			AccountID:    cloudflare.F("account_id"),
			ID:           cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Account:      cloudflare.F("account"),
			CustIP:       cloudflare.F("192.168.3.4/31"),
			Interconnect: cloudflare.F("interconnect"),
			Magic: cloudflare.F(network_interconnects.CNIUpdateParamsMagic{
				ConduitName: cloudflare.F("conduit_name"),
				Description: cloudflare.F("description"),
				Mtu:         cloudflare.F(int64(0)),
			}),
			P2pIP: cloudflare.F("192.168.3.4/31"),
			BGP: cloudflare.F(network_interconnects.CNIUpdateParamsBGP{
				CustomerASN:   cloudflare.F(int64(0)),
				ExtraPrefixes: cloudflare.F([]string{"string"}),
				Md5Key:        cloudflare.F("md5_key"),
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

func TestCNIListWithOptionalParams(t *testing.T) {
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
	_, err := client.NetworkInterconnects.CNIs.List(context.TODO(), network_interconnects.CNIListParams{
		AccountID: cloudflare.F("account_id"),
		Cursor:    cloudflare.F(int64(0)),
		Limit:     cloudflare.F(int64(0)),
		Slot:      cloudflare.F("slot"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCNIDelete(t *testing.T) {
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
	err := client.NetworkInterconnects.CNIs.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		network_interconnects.CNIDeleteParams{
			AccountID: cloudflare.F("account_id"),
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

func TestCNIGet(t *testing.T) {
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
	_, err := client.NetworkInterconnects.CNIs.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		network_interconnects.CNIGetParams{
			AccountID: cloudflare.F("account_id"),
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

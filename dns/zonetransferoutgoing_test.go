// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestZoneTransferOutgoingNew(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.New(context.TODO(), dns.ZoneTransferOutgoingNewParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
		Name:   cloudflare.F("www.example.com."),
		Peers:  cloudflare.F([]string{"23ff594956f20c2a721606e94745a8aa", "00920f38ce07c2e2f4df50b1f61d4194"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingUpdate(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.Update(context.TODO(), dns.ZoneTransferOutgoingUpdateParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
		Name:   cloudflare.F("www.example.com."),
		Peers:  cloudflare.F([]string{"23ff594956f20c2a721606e94745a8aa", "00920f38ce07c2e2f4df50b1f61d4194"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingDelete(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.Delete(context.TODO(), dns.ZoneTransferOutgoingDeleteParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingDisable(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.Disable(context.TODO(), dns.ZoneTransferOutgoingDisableParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
		Body:   map[string]interface{}{},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingEnable(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.Enable(context.TODO(), dns.ZoneTransferOutgoingEnableParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
		Body:   map[string]interface{}{},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingForceNotify(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.ForceNotify(context.TODO(), dns.ZoneTransferOutgoingForceNotifyParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
		Body:   map[string]interface{}{},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneTransferOutgoingGet(t *testing.T) {
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
	_, err := client.DNS.ZoneTransfers.Outgoing.Get(context.TODO(), dns.ZoneTransferOutgoingGetParams{
		ZoneID: cloudflare.F("269d8f4853475ca241c4e730be286b20"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
)

func TestGatewayListItemList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Gateway.Lists.Items.List(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayListItemListParams{
			AccountID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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

func TestGatewayListItemListAutoPaging(t *testing.T) {
	var page1, page2 int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("page") {
		case "", "1":
			page1++
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"result":[{"created_at":"2024-01-01T00:00:00Z","description":"one","value":"one"}],"result_info":{"page":1,"per_page":50}}`))
		case "2":
			page2++
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"result":[{"created_at":"2024-01-01T00:00:00Z","description":"two","value":"two"}],"result_info":{"page":2,"per_page":50}}`))
		case "3":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"result":[],"result_info":{"page":3,"per_page":50}}`))
		default:
			http.Error(w, "unexpected page", http.StatusBadRequest)
		}
	}))
	defer server.Close()

	client := cloudflare.NewClient(
		option.WithBaseURL(server.URL),
		option.WithAPIKey("test"),
		option.WithAPIEmail("user@example.com"),
	)

	pager := client.ZeroTrust.Gateway.Lists.Items.ListAutoPaging(
		context.Background(),
		"list-123",
		zero_trust.GatewayListItemListParams{AccountID: cloudflare.F("acct-123")},
	)

	var got []string
	for pager.Next() {
		got = append(got, pager.Current().Value)
	}
	if err := pager.Err(); err != nil {
		t.Fatalf("pager.Err() = %v", err)
	}
	if len(got) != 2 || got[0] != "one" || got[1] != "two" {
		t.Fatalf("pager returned %v, want [one two]", got)
	}
	if page1 != 1 || page2 != 1 {
		t.Fatalf("pages requested = (%d, %d), want (1, 1)", page1, page2)
	}
}

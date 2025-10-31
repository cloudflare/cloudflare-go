// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package servers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/servers"
)

func TestAccessAIControlMcpServerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.New(context.TODO(), servers.AccessAIControlMcpServerNewParams{
		AccountID:       cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
		ID:              cloudflare.F("my-mcp-server"),
		AuthType:        cloudflare.F(servers.AccessAIControlMcpServerNewParamsAuthTypeUnauthenticated),
		Hostname:        cloudflare.F("https://exmaple.com/mcp"),
		Name:            cloudflare.F("My MCP Server"),
		AuthCredentials: cloudflare.F("auth_credentials"),
		Description:     cloudflare.F("This is one remote mcp server"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessAIControlMcpServerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Update(
		context.TODO(),
		"my-mcp-server",
		servers.AccessAIControlMcpServerUpdateParams{
			AccountID:       cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
			AuthCredentials: cloudflare.F("auth_credentials"),
			Description:     cloudflare.F("This is one remote mcp server"),
			Name:            cloudflare.F("My MCP Server"),
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

func TestAccessAIControlMcpServerListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.List(context.TODO(), servers.AccessAIControlMcpServerListParams{
		AccountID: cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
		Search:    cloudflare.F("search"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessAIControlMcpServerDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Delete(
		context.TODO(),
		"my-mcp-server",
		servers.AccessAIControlMcpServerDeleteParams{
			AccountID: cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
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

func TestAccessAIControlMcpServerRead(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Read(
		context.TODO(),
		"my-mcp-server",
		servers.AccessAIControlMcpServerReadParams{
			AccountID: cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
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

func TestAccessAIControlMcpServerSync(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Sync(
		context.TODO(),
		"my-mcp-portal",
		servers.AccessAIControlMcpServerSyncParams{
			AccountID: cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
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

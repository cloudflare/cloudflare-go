// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
)

func TestAccessAIControlMcpPortalNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.New(context.TODO(), zero_trust.AccessAIControlMcpPortalNewParams{
		AccountID:        cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
		ID:               cloudflare.F("my-mcp-portal"),
		Hostname:         cloudflare.F("exmaple.com"),
		Name:             cloudflare.F("My MCP Portal"),
		Description:      cloudflare.F("This is my custom MCP Portal"),
		SecureWebGateway: cloudflare.F(false),
		Servers: cloudflare.F([]zero_trust.AccessAIControlMcpPortalNewParamsServer{{
			ServerID:        cloudflare.F("my-mcp-server"),
			DefaultDisabled: cloudflare.F(true),
			OnBehalf:        cloudflare.F(true),
			UpdatedPrompts: cloudflare.F([]zero_trust.AccessAIControlMcpPortalNewParamsServersUpdatedPrompt{{
				Name:        cloudflare.F("name"),
				Description: cloudflare.F("description"),
				Enabled:     cloudflare.F(true),
			}}),
			UpdatedTools: cloudflare.F([]zero_trust.AccessAIControlMcpPortalNewParamsServersUpdatedTool{{
				Name:        cloudflare.F("name"),
				Description: cloudflare.F("description"),
				Enabled:     cloudflare.F(true),
			}}),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessAIControlMcpPortalUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.Update(
		context.TODO(),
		"my-mcp-portal",
		zero_trust.AccessAIControlMcpPortalUpdateParams{
			AccountID:        cloudflare.F("a86a8f5c339544d7bdc89926de14fb8c"),
			Description:      cloudflare.F("This is my custom MCP Portal"),
			Hostname:         cloudflare.F("exmaple.com"),
			Name:             cloudflare.F("My MCP Portal"),
			SecureWebGateway: cloudflare.F(false),
			Servers: cloudflare.F([]zero_trust.AccessAIControlMcpPortalUpdateParamsServer{{
				ServerID:        cloudflare.F("my-mcp-server"),
				DefaultDisabled: cloudflare.F(true),
				OnBehalf:        cloudflare.F(true),
				UpdatedPrompts: cloudflare.F([]zero_trust.AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt{{
					Name:        cloudflare.F("name"),
					Description: cloudflare.F("description"),
					Enabled:     cloudflare.F(true),
				}}),
				UpdatedTools: cloudflare.F([]zero_trust.AccessAIControlMcpPortalUpdateParamsServersUpdatedTool{{
					Name:        cloudflare.F("name"),
					Description: cloudflare.F("description"),
					Enabled:     cloudflare.F(true),
				}}),
			}}),
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

func TestAccessAIControlMcpPortalListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.List(context.TODO(), zero_trust.AccessAIControlMcpPortalListParams{
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

func TestAccessAIControlMcpPortalDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.Delete(
		context.TODO(),
		"my-mcp-portal",
		zero_trust.AccessAIControlMcpPortalDeleteParams{
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

func TestAccessAIControlMcpPortalRead(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.Read(
		context.TODO(),
		"my-mcp-portal",
		zero_trust.AccessAIControlMcpPortalReadParams{
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

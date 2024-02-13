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

func TestTeamnetRouteTunnelRouteListTunnelRoutesWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Teamnets.Routes.TunnelRouteListTunnelRoutes(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.TeamnetRouteTunnelRouteListTunnelRoutesParams{
			Comment:          cloudflare.F("Example comment for this route."),
			ExistedAt:        cloudflare.F[any](map[string]interface{}{}),
			IsDeleted:        cloudflare.F[any](map[string]interface{}{}),
			NetworkSubset:    cloudflare.F[any](map[string]interface{}{}),
			NetworkSuperset:  cloudflare.F[any](map[string]interface{}{}),
			Page:             cloudflare.F(1.000000),
			PerPage:          cloudflare.F(1.000000),
			TunTypes:         cloudflare.F("cfd_tunnel,warp_connector"),
			TunnelID:         cloudflare.F[any](map[string]interface{}{}),
			VirtualNetworkID: cloudflare.F[any](map[string]interface{}{}),
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

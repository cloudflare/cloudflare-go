// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
)

func TestAccessInfrastructureTargetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Infrastructure.Targets.New(context.TODO(), zero_trust.AccessInfrastructureTargetNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Hostname:  cloudflare.F("infra-access-target"),
		IP: cloudflare.F(zero_trust.AccessInfrastructureTargetNewParamsIP{
			IPV4: cloudflare.F(zero_trust.AccessInfrastructureTargetNewParamsIPIPV4{
				IPAddr:           cloudflare.F("187.26.29.249"),
				VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
			}),
			IPV6: cloudflare.F(zero_trust.AccessInfrastructureTargetNewParamsIPIPV6{
				IPAddr:           cloudflare.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
				VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
			}),
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

func TestAccessInfrastructureTargetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Infrastructure.Targets.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.AccessInfrastructureTargetUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Hostname:  cloudflare.F("infra-access-target"),
			IP: cloudflare.F(zero_trust.AccessInfrastructureTargetUpdateParamsIP{
				IPV4: cloudflare.F(zero_trust.AccessInfrastructureTargetUpdateParamsIPIPV4{
					IPAddr:           cloudflare.F("187.26.29.249"),
					VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
				}),
				IPV6: cloudflare.F(zero_trust.AccessInfrastructureTargetUpdateParamsIPIPV6{
					IPAddr:           cloudflare.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
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

func TestAccessInfrastructureTargetListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Infrastructure.Targets.List(context.TODO(), zero_trust.AccessInfrastructureTargetListParams{
		AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		CreatedAfter:     cloudflare.F(time.Now()),
		CreatedBefore:    cloudflare.F(time.Now()),
		Direction:        cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsDirectionAsc),
		Hostname:         cloudflare.F("hostname"),
		HostnameContains: cloudflare.F("hostname_contains"),
		IPV4:             cloudflare.F("ip_v4"),
		IPV6:             cloudflare.F("ip_v6"),
		IPs:              cloudflare.F([]string{"string"}),
		ModifiedAfter:    cloudflare.F(time.Now()),
		ModifiedBefore:   cloudflare.F(time.Now()),
		Order:            cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsOrderHostname),
		Page:             cloudflare.F(int64(1)),
		PerPage:          cloudflare.F(int64(1)),
		VirtualNetworkID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessInfrastructureTargetDelete(t *testing.T) {
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
	err := client.ZeroTrust.Access.Infrastructure.Targets.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.AccessInfrastructureTargetDeleteParams{
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

func TestAccessInfrastructureTargetBulkDelete(t *testing.T) {
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
	err := client.ZeroTrust.Access.Infrastructure.Targets.BulkDelete(context.TODO(), zero_trust.AccessInfrastructureTargetBulkDeleteParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessInfrastructureTargetBulkUpdate(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Infrastructure.Targets.BulkUpdate(context.TODO(), zero_trust.AccessInfrastructureTargetBulkUpdateParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body: []zero_trust.AccessInfrastructureTargetBulkUpdateParamsBody{{
			Hostname: cloudflare.F("infra-access-target"),
			IP: cloudflare.F(zero_trust.AccessInfrastructureTargetBulkUpdateParamsBodyIP{
				IPV4: cloudflare.F(zero_trust.AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4{
					IPAddr:           cloudflare.F("187.26.29.249"),
					VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
				}),
				IPV6: cloudflare.F(zero_trust.AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6{
					IPAddr:           cloudflare.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					VirtualNetworkID: cloudflare.F("c77b744e-acc8-428f-9257-6878c046ed55"),
				}),
			}),
		}},
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessInfrastructureTargetGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Infrastructure.Targets.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		zero_trust.AccessInfrastructureTargetGetParams{
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package k2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/k2"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestStreamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.K2.Streams.New(context.TODO(), k2.StreamNewParams{
		AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("my_k2_stream"),
		HTTP: cloudflare.F(k2.StreamNewParamsHTTP{
			Enabled:        cloudflare.F(true),
			Authentication: cloudflare.F(true),
			CORS: cloudflare.F(k2.StreamNewParamsHTTPCORS{
				Origins: cloudflare.F([]string{"string"}),
			}),
		}),
		RetentionSeconds: cloudflare.F(int64(3600)),
		WorkerBinding: cloudflare.F[k2.StreamNewParamsWorkerBindingUnion](k2.StreamNewParamsWorkerBindingEnabled{
			Enabled: cloudflare.F(k2.StreamNewParamsWorkerBindingEnabledEnabledFalse),
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

func TestStreamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.K2.Streams.Update(
		context.TODO(),
		"053e105f4ecef8ad9ca31a8372d0c353",
		k2.StreamUpdateParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
			HTTP: cloudflare.F(k2.StreamUpdateParamsHTTP{
				Enabled:        cloudflare.F(true),
				Authentication: cloudflare.F(true),
				CORS: cloudflare.F(k2.StreamUpdateParamsHTTPCORS{
					Origins: cloudflare.F([]string{"string"}),
				}),
			}),
			RetentionSeconds: cloudflare.F(int64(3600)),
			WorkerBinding: cloudflare.F[k2.StreamUpdateParamsWorkerBindingUnion](k2.StreamUpdateParamsWorkerBindingEnabled{
				Enabled: cloudflare.F(k2.StreamUpdateParamsWorkerBindingEnabledEnabledFalse),
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

func TestStreamListWithOptionalParams(t *testing.T) {
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
	_, err := client.K2.Streams.List(context.TODO(), k2.StreamListParams{
		AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
		Name:      cloudflare.F("x"),
		Page:      cloudflare.F(int64(1)),
		PerPage:   cloudflare.F(int64(1)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStreamDelete(t *testing.T) {
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
	_, err := client.K2.Streams.Delete(
		context.TODO(),
		"053e105f4ecef8ad9ca31a8372d0c353",
		k2.StreamDeleteParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
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

func TestStreamGet(t *testing.T) {
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
	_, err := client.K2.Streams.Get(
		context.TODO(),
		"053e105f4ecef8ad9ca31a8372d0c353",
		k2.StreamGetParams{
			AccountID: cloudflare.F("0123105f4ecef8ad9ca31a8372d0c353"),
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

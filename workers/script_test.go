// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/workers"
)

func TestScriptUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Workers.Scripts.Update(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body: workers.ScriptUpdateParamsBodyMetadata{
				Metadata: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadata{
					Assets: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataAssets{
						Config: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataAssetsConfig{
							HTMLHandling:     cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
							NotFoundHandling: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone),
							ServeDirectly:    cloudflare.F(true),
						}),
						JWT: cloudflare.F("jwt"),
					}),
					Bindings: cloudflare.F([]workers.ScriptUpdateParamsBodyMetadataMetadataBinding{{
						Name: cloudflare.F("MY_ENV_VAR"),
						Type: cloudflare.F("plain_text"),
					}}),
					BodyPart:           cloudflare.F("worker.js"),
					CompatibilityDate:  cloudflare.F("2021-01-01"),
					CompatibilityFlags: cloudflare.F([]string{"nodejs_compat"}),
					KeepAssets:         cloudflare.F(false),
					KeepBindings:       cloudflare.F([]string{"string"}),
					Logpush:            cloudflare.F(false),
					MainModule:         cloudflare.F("worker.js"),
					Migrations: cloudflare.F[workers.ScriptUpdateParamsBodyMetadataMetadataMigrationsUnion](workers.SingleStepMigrationParam{
						DeletedClasses:   cloudflare.F([]string{"string"}),
						NewClasses:       cloudflare.F([]string{"string"}),
						NewSqliteClasses: cloudflare.F([]string{"string"}),
						NewTag:           cloudflare.F("v2"),
						OldTag:           cloudflare.F("v1"),
						RenamedClasses: cloudflare.F([]workers.SingleStepMigrationRenamedClassParam{{
							From: cloudflare.F("from"),
							To:   cloudflare.F("to"),
						}}),
						TransferredClasses: cloudflare.F([]workers.SingleStepMigrationTransferredClassParam{{
							From:       cloudflare.F("from"),
							FromScript: cloudflare.F("from_script"),
							To:         cloudflare.F("to"),
						}}),
					}),
					Observability: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataObservability{
						Enabled:          cloudflare.F(true),
						HeadSamplingRate: cloudflare.F(0.100000),
					}),
					Placement: cloudflare.F(workers.PlacementConfigurationParam{
						Mode: cloudflare.F(workers.PlacementConfigurationModeSmart),
					}),
					Tags: cloudflare.F([]string{"string"}),
					TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
						Service:     cloudflare.F("my-log-consumer"),
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
					}}),
					UsageModel: cloudflare.F(workers.ScriptUpdateParamsBodyMetadataMetadataUsageModelBundled),
					VersionTags: cloudflare.F(map[string]string{
						"foo": "string",
					}),
				}),
			},
			RollbackTo: cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
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

func TestScriptList(t *testing.T) {
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
	_, err := client.Workers.Scripts.List(context.TODO(), workers.ScriptListParams{
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

func TestScriptDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Workers.Scripts.Delete(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Force:     cloudflare.F(true),
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

func TestScriptGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	resp, err := client.Workers.Scripts.Get(
		context.TODO(),
		"this-is_my_script-01",
		workers.ScriptGetParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

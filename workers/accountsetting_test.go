// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
  "context"
  "errors"
  "os"
  "testing"

  "github.com/cloudflare/cloudflare-go/v2"
  "github.com/cloudflare/cloudflare-go/v2/internal/testutil"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/workers"
)

func TestAccountSettingUpdateWithOptionalParams(t *testing.T) {
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
  _, err := client.Workers.AccountSettings.Update(context.TODO(), workers.AccountSettingUpdateParams{
    AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
    DefaultUsageModel: cloudflare.F("default_usage_model"),
    GreenCompute: cloudflare.F(true),
  })
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestAccountSettingGet(t *testing.T) {
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
  _, err := client.Workers.AccountSettings.Get(context.TODO(), workers.AccountSettingGetParams{
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

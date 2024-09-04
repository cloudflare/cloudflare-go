// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
  "context"
  "errors"
  "os"
  "testing"
  "time"

  "github.com/cloudflare/cloudflare-go/v2"
  "github.com/cloudflare/cloudflare-go/v2/internal/testutil"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestVerifiedBotTopBotsWithOptionalParams(t *testing.T) {
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
  _, err := client.Radar.VerifiedBots.Top.Bots(context.TODO(), radar.VerifiedBotTopBotsParams{
    ASN: cloudflare.F([]string{"string", "string", "string"}),
    Continent: cloudflare.F([]string{"string", "string", "string"}),
    DateEnd: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
    DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    Format: cloudflare.F(radar.VerifiedBotTopBotsParamsFormatJson),
    Limit: cloudflare.F(int64(5)),
    Location: cloudflare.F([]string{"string", "string", "string"}),
    Name: cloudflare.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestVerifiedBotTopCategoriesWithOptionalParams(t *testing.T) {
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
  _, err := client.Radar.VerifiedBots.Top.Categories(context.TODO(), radar.VerifiedBotTopCategoriesParams{
    ASN: cloudflare.F([]string{"string", "string", "string"}),
    Continent: cloudflare.F([]string{"string", "string", "string"}),
    DateEnd: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
    DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    Format: cloudflare.F(radar.VerifiedBotTopCategoriesParamsFormatJson),
    Limit: cloudflare.F(int64(5)),
    Location: cloudflare.F([]string{"string", "string", "string"}),
    Name: cloudflare.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

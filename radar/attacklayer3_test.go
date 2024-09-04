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

func TestAttackLayer3TimeseriesWithOptionalParams(t *testing.T) {
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
  _, err := client.Radar.Attacks.Layer3.Timeseries(context.TODO(), radar.AttackLayer3TimeseriesParams{
    AggInterval: cloudflare.F(radar.AttackLayer3TimeseriesParamsAggInterval15m),
    ASN: cloudflare.F([]string{"string", "string", "string"}),
    Continent: cloudflare.F([]string{"string", "string", "string"}),
    DateEnd: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    DateRange: cloudflare.F([]string{"7d", "7d", "7d"}),
    DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
    Direction: cloudflare.F(radar.AttackLayer3TimeseriesParamsDirectionOrigin),
    Format: cloudflare.F(radar.AttackLayer3TimeseriesParamsFormatJson),
    IPVersion: cloudflare.F([]radar.AttackLayer3TimeseriesParamsIPVersion{radar.AttackLayer3TimeseriesParamsIPVersionIPv4, radar.AttackLayer3TimeseriesParamsIPVersionIPv6}),
    Location: cloudflare.F([]string{"string", "string", "string"}),
    Metric: cloudflare.F(radar.AttackLayer3TimeseriesParamsMetricBytes),
    Name: cloudflare.F([]string{"string", "string", "string"}),
    Normalization: cloudflare.F(radar.AttackLayer3TimeseriesParamsNormalizationPercentageChange),
    Protocol: cloudflare.F([]radar.AttackLayer3TimeseriesParamsProtocol{radar.AttackLayer3TimeseriesParamsProtocolUdp, radar.AttackLayer3TimeseriesParamsProtocolTCP, radar.AttackLayer3TimeseriesParamsProtocolIcmp}),
  })
  if err != nil {
    var apierr *cloudflare.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

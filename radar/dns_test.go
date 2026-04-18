// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/radar"
)

func TestDNSTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Timeseries(context.TODO(), radar.DNSTimeseriesParams{
		AggInterval:    cloudflare.F(radar.DNSTimeseriesParamsAggInterval1h),
		ASN:            cloudflare.F([]string{"string"}),
		CacheHit:       cloudflare.F([]bool{true}),
		Continent:      cloudflare.F([]string{"string"}),
		DateEnd:        cloudflare.F([]time.Time{time.Now()}),
		DateRange:      cloudflare.F([]string{"7d"}),
		DateStart:      cloudflare.F([]time.Time{time.Now()}),
		DNSSEC:         cloudflare.F([]radar.DNSTimeseriesParamsDNSSEC{radar.DNSTimeseriesParamsDNSSECInvalid}),
		DNSSECAware:    cloudflare.F([]radar.DNSTimeseriesParamsDNSSECAware{radar.DNSTimeseriesParamsDNSSECAwareSupported}),
		DNSSECE2E:      cloudflare.F([]bool{true}),
		Format:         cloudflare.F(radar.DNSTimeseriesParamsFormatJson),
		IPVersion:      cloudflare.F([]radar.DNSTimeseriesParamsIPVersion{radar.DNSTimeseriesParamsIPVersionIPv4}),
		Location:       cloudflare.F([]string{"string"}),
		MatchingAnswer: cloudflare.F([]bool{true}),
		Name:           cloudflare.F([]string{"main_series"}),
		Nodata:         cloudflare.F([]bool{true}),
		Protocol:       cloudflare.F([]radar.DNSTimeseriesParamsProtocol{radar.DNSTimeseriesParamsProtocolUdp}),
		QueryType:      cloudflare.F([]radar.DNSTimeseriesParamsQueryType{radar.DNSTimeseriesParamsQueryTypeA}),
		ResponseCode:   cloudflare.F([]radar.DNSTimeseriesParamsResponseCode{radar.DNSTimeseriesParamsResponseCodeNoerror}),
		ResponseTTL:    cloudflare.F([]radar.DNSTimeseriesParamsResponseTTL{radar.DNSTimeseriesParamsResponseTTLLte1M}),
		Tld:            cloudflare.F([]string{"com"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

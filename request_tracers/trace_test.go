// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/request_tracers"
)

func TestTraceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RequestTracers.Traces.New(context.TODO(), request_tracers.TraceNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Method:    cloudflare.F("PUT"),
		URL:       cloudflare.F("https://some.zone/some_path"),
		Body: cloudflare.F(request_tracers.TraceNewParamsBody{
			Base64:    cloudflare.F("c29tZV9yZXF1ZXN0X2JvZHk="),
			Json:      cloudflare.F[any](map[string]interface{}{}),
			PlainText: cloudflare.F("plain_text"),
		}),
		Context: cloudflare.F(request_tracers.TraceNewParamsContext{
			BotScore: cloudflare.F(int64(0)),
			Geoloc: cloudflare.F(request_tracers.TraceNewParamsContextGeoloc{
				City:                cloudflare.F("London"),
				Continent:           cloudflare.F("continent"),
				IsEuCountry:         cloudflare.F(true),
				ISOCode:             cloudflare.F("iso_code"),
				Latitude:            cloudflare.F(0.000000),
				Longitude:           cloudflare.F(0.000000),
				PostalCode:          cloudflare.F("postal_code"),
				RegionCode:          cloudflare.F("region_code"),
				Subdivision2ISOCode: cloudflare.F("subdivision_2_iso_code"),
				Timezone:            cloudflare.F("timezone"),
			}),
			SkipChallenge: cloudflare.F(true),
			ThreatScore:   cloudflare.F(int64(0)),
		}),
		Cookies: cloudflare.F(map[string]string{
			"cookie_name_1": "cookie_value_1",
			"cookie_name_2": "cookie_value_2",
		}),
		Headers: cloudflare.F(map[string]string{
			"header_name_1": "header_value_1",
			"header_name_2": "header_value_2",
		}),
		Protocol:     cloudflare.F("HTTP/1.1"),
		SkipResponse: cloudflare.F(true),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

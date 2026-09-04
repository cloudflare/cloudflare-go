// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/radar"
)

func TestAnnotationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Annotations.List(context.TODO(), radar.AnnotationListParams{
		ASN:         cloudflare.F(int64(174)),
		Bot:         cloudflare.F("gptbot"),
		CA:          cloudflare.F("ca"),
		DataSource:  cloudflare.F(radar.AnnotationListParamsDataSourceAll),
		DateEnd:     cloudflare.F(time.Now()),
		DateRange:   cloudflare.F("7d"),
		DateStart:   cloudflare.F(time.Now()),
		EventType:   cloudflare.F(radar.AnnotationListParamsEventTypeOutage),
		Format:      cloudflare.F(radar.AnnotationListParamsFormatJson),
		GeoID:       cloudflare.F("3190509"),
		Limit:       cloudflare.F(int64(1)),
		Location:    cloudflare.F("US"),
		Log:         cloudflare.F("log"),
		Offset:      cloudflare.F(int64(0)),
		Origin:      cloudflare.F("amazon-us-east-1"),
		OutageCause: cloudflare.F(radar.AnnotationListParamsOutageCauseBlocking),
		OutageType:  cloudflare.F(radar.AnnotationListParamsOutageTypeNationwide),
		Query:       cloudflare.F("portugal"),
		Tags:        cloudflare.F([]radar.AnnotationListParamsTag{radar.AnnotationListParamsTagAdm1}),
		TLD:         cloudflare.F("com"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

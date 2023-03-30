package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/core"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
)

func TestZonesNewWithOptionalParams(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.New(context.TODO(), &requests.ZoneNewParams{Account: fields.F(requests.ZoneNewParamsAccount{ID: fields.F("023e105f4ecef8ad9ca31a8372d0c353")}), Name: fields.F("example.com"), Type: fields.F(requests.TypeFo5LnpAtFull)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonesGet(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonesUpdateWithOptionalParams(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		&requests.ZoneUpdateParams{Plan: fields.F(requests.ZoneUpdateParamsPlan{ID: fields.F("023e105f4ecef8ad9ca31a8372d0c353")}), Type: fields.F(requests.ZoneUpdateParamsTypeFull), VanityNameServers: fields.F([]string{"string", "string", "string"})},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonesListWithOptionalParams(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.List(context.TODO(), &requests.ZoneListParams{Name: fields.F("example.com"), Status: fields.F("string"), Page: fields.F(1.000000), PerPage: fields.F(5.000000), Order: fields.F(requests.ZoneListParamsOrderName), Direction: fields.F(requests.ZoneListParamsDirectionAsc), Match: fields.F(requests.ZoneListParamsMatchAny), Account: fields.F(requests.ZoneListParamsAccount{ID: fields.F("string"), Name: fields.F("string")})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZonesDelete(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

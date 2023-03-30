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

func TestImageResizingsUpdate(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Settings.ImageResizings.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		&requests.ImageResizingUpdateParams{Value: fields.F(requests.ImageResizing{ID: fields.F(requests.ImageResizingIDImageResizing), Value: fields.F(requests.ImageResizingValueOn)})},
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

func TestImageResizingsList(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Settings.ImageResizings.List(
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

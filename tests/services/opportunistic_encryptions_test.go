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

func TestOpportunisticEncryptionsUpdate(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Settings.OpportunisticEncryptions.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		&requests.OpportunisticEncryptionUpdateParams{Value: fields.F(requests.OpportunisticEncryptionValueOn)},
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

func TestOpportunisticEncryptionsList(t *testing.T) {
	c := cloudflare.NewCloudflare(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Zones.Settings.OpportunisticEncryptions.List(
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

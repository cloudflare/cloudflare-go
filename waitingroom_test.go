// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestWaitingRoomNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.WaitingRoomNewParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.WaitingRoomNewParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.WaitingRoomNewParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.WaitingRoomNewParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.WaitingRoomNewParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.WaitingRoomNewParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.WaitingRoomNewParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.WaitingRoomNewParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestWaitingRoomUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.WaitingRoomUpdateParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.WaitingRoomUpdateParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.WaitingRoomUpdateParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.WaitingRoomUpdateParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.WaitingRoomUpdateParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.WaitingRoomUpdateParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.WaitingRoomUpdateParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.WaitingRoomUpdateParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestWaitingRoomList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.WaitingRoomEditParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.WaitingRoomEditParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.WaitingRoomEditParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.WaitingRoomEditParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.WaitingRoomEditParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.WaitingRoomEditParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.WaitingRoomEditParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.WaitingRoomEditParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestWaitingRoomGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomPreview(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.WaitingRooms.Preview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.WaitingRoomPreviewParams{
			CustomHTML: cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
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

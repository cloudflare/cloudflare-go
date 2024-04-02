// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
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
	)
	_, err := client.WaitingRooms.New(context.TODO(), waiting_rooms.WaitingRoomNewParams{
		ZoneID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Host:              cloudflare.F("shop.example.com"),
		Name:              cloudflare.F("production_webinar"),
		NewUsersPerMinute: cloudflare.F(int64(200)),
		TotalActiveUsers:  cloudflare.F(int64(200)),
		AdditionalRoutes: cloudflare.F([]waiting_rooms.WaitingRoomNewParamsAdditionalRoute{{
			Host: cloudflare.F("shop2.example.com"),
			Path: cloudflare.F("/shop2/checkout"),
		}, {
			Host: cloudflare.F("shop2.example.com"),
			Path: cloudflare.F("/shop2/checkout"),
		}, {
			Host: cloudflare.F("shop2.example.com"),
			Path: cloudflare.F("/shop2/checkout"),
		}}),
		CookieAttributes: cloudflare.F(waiting_rooms.WaitingRoomNewParamsCookieAttributes{
			Samesite: cloudflare.F(waiting_rooms.WaitingRoomNewParamsCookieAttributesSamesiteAuto),
			Secure:   cloudflare.F(waiting_rooms.WaitingRoomNewParamsCookieAttributesSecureAuto),
		}),
		CookieSuffix:            cloudflare.F("abcd"),
		CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
		DefaultTemplateLanguage: cloudflare.F(waiting_rooms.WaitingRoomNewParamsDefaultTemplateLanguageEsEs),
		Description:             cloudflare.F("Production - DO NOT MODIFY"),
		DisableSessionRenewal:   cloudflare.F(false),
		JsonResponseEnabled:     cloudflare.F(false),
		Path:                    cloudflare.F("/shop/checkout"),
		QueueAll:                cloudflare.F(true),
		QueueingMethod:          cloudflare.F(waiting_rooms.WaitingRoomNewParamsQueueingMethodFifo),
		QueueingStatusCode:      cloudflare.F(waiting_rooms.WaitingRoomNewParamsQueueingStatusCode202),
		SessionDuration:         cloudflare.F(int64(1)),
		Suspended:               cloudflare.F(true),
	})
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
	)
	_, err := client.WaitingRooms.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomUpdateParams{
			ZoneID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]waiting_rooms.WaitingRoomUpdateParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsCookieAttributes{
				Samesite: cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(waiting_rooms.WaitingRoomUpdateParamsQueueingStatusCode202),
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
	)
	_, err := client.WaitingRooms.List(context.TODO(), waiting_rooms.WaitingRoomListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
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
	)
	_, err := client.WaitingRooms.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
	)
	_, err := client.WaitingRooms.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomEditParams{
			ZoneID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]waiting_rooms.WaitingRoomEditParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(waiting_rooms.WaitingRoomEditParamsCookieAttributes{
				Samesite: cloudflare.F(waiting_rooms.WaitingRoomEditParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(waiting_rooms.WaitingRoomEditParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(waiting_rooms.WaitingRoomEditParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(waiting_rooms.WaitingRoomEditParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(waiting_rooms.WaitingRoomEditParamsQueueingStatusCode202),
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
	)
	_, err := client.WaitingRooms.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/waiting_rooms"
)

func TestWaitingRoomNewWithOptionalParams(t *testing.T) {
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
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Query: waiting_rooms.QueryParam{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
				Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			EnabledOriginCommands:   cloudflare.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
		},
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
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Query: waiting_rooms.QueryParam{
				Host:              cloudflare.F("shop.example.com"),
				Name:              cloudflare.F("production_webinar"),
				NewUsersPerMinute: cloudflare.F(int64(200)),
				TotalActiveUsers:  cloudflare.F(int64(200)),
				AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}}),
				CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
					Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cloudflare.F("abcd"),
				CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             cloudflare.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cloudflare.F(false),
				EnabledOriginCommands:   cloudflare.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
				JsonResponseEnabled:     cloudflare.F(false),
				Path:                    cloudflare.F("/shop/checkout"),
				QueueAll:                cloudflare.F(true),
				QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         cloudflare.F(int64(1)),
				Suspended:               cloudflare.F(true),
			},
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

func TestWaitingRoomListWithOptionalParams(t *testing.T) {
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
		ZoneID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Page:    cloudflare.F(1.000000),
		PerPage: cloudflare.F(5.000000),
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
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Query: waiting_rooms.QueryParam{
				Host:              cloudflare.F("shop.example.com"),
				Name:              cloudflare.F("production_webinar"),
				NewUsersPerMinute: cloudflare.F(int64(200)),
				TotalActiveUsers:  cloudflare.F(int64(200)),
				AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}}),
				CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
					Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cloudflare.F("abcd"),
				CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             cloudflare.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cloudflare.F(false),
				EnabledOriginCommands:   cloudflare.F([]waiting_rooms.QueryEnabledOriginCommand{waiting_rooms.QueryEnabledOriginCommandRevoke}),
				JsonResponseEnabled:     cloudflare.F(false),
				Path:                    cloudflare.F("/shop/checkout"),
				QueueAll:                cloudflare.F(true),
				QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         cloudflare.F(int64(1)),
				Suspended:               cloudflare.F(true),
			},
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

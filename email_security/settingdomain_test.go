// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/cloudflare-go/v7/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

func TestSettingDomainNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.New(context.TODO(), email_security.SettingDomainNewParams{
		AccountID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainNewParamsAllowedDeliveryMode{email_security.SettingDomainNewParamsAllowedDeliveryModeDirect}),
		Domain:               cloudflare.F("domain"),
		DropDispositions:     cloudflare.F([]email_security.SettingDomainNewParamsDropDisposition{email_security.SettingDomainNewParamsDropDispositionMalicious}),
		IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
		Regions:              cloudflare.F([]email_security.SettingDomainNewParamsRegion{email_security.SettingDomainNewParamsRegionGlobal}),
		Folder:               cloudflare.F(email_security.SettingDomainNewParamsFolderAllItems),
		IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		LookbackHops:         cloudflare.F(int64(1)),
		RequireTLSInbound:    cloudflare.F(true),
		RequireTLSOutbound:   cloudflare.F(true),
		Transport:            cloudflare.F("transport"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingDomainUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingDomainUpdateParams{
			AccountID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainUpdateParamsAllowedDeliveryMode{email_security.SettingDomainUpdateParamsAllowedDeliveryModeDirect}),
			DropDispositions:     cloudflare.F([]email_security.SettingDomainUpdateParamsDropDisposition{email_security.SettingDomainUpdateParamsDropDispositionMalicious}),
			IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			Regions:              cloudflare.F([]email_security.SettingDomainUpdateParamsRegion{email_security.SettingDomainUpdateParamsRegionGlobal}),
			Folder:               cloudflare.F(email_security.SettingDomainUpdateParamsFolderAllItems),
			IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LookbackHops:         cloudflare.F(int64(1)),
			RequireTLSInbound:    cloudflare.F(true),
			RequireTLSOutbound:   cloudflare.F(true),
			Transport:            cloudflare.F("transport"),
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

func TestSettingDomainListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.List(context.TODO(), email_security.SettingDomainListParams{
		AccountID:           cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ActiveDeliveryMode:  cloudflare.F(email_security.SettingDomainListParamsActiveDeliveryModeDirect),
		AllowedDeliveryMode: cloudflare.F(email_security.SettingDomainListParamsAllowedDeliveryModeDirect),
		Direction:           cloudflare.F(email_security.SettingDomainListParamsDirectionAsc),
		Domain:              cloudflare.F([]string{"string"}),
		IntegrationID:       cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Order:               cloudflare.F(email_security.SettingDomainListParamsOrderDomain),
		Page:                cloudflare.F(int64(1)),
		PerPage:             cloudflare.F(int64(20)),
		Search:              cloudflare.F("search"),
		Status:              cloudflare.F(email_security.SettingDomainListParamsStatusPending),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingDomainDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingDomainDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

func TestSettingDomainBatch(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.Batch(context.TODO(), email_security.SettingDomainBatchParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Deletes: cloudflare.F([]email_security.SettingDomainBatchParamsDelete{{
			ID: cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
		}}),
		Patches: cloudflare.F([]email_security.SettingDomainBatchParamsPatch{{
			ID:                   cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
			AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainBatchParamsPatchesAllowedDeliveryMode{email_security.SettingDomainBatchParamsPatchesAllowedDeliveryModeDirect}),
			DropDispositions:     cloudflare.F([]email_security.SettingDomainBatchParamsPatchesDropDisposition{email_security.SettingDomainBatchParamsPatchesDropDispositionMalicious}),
			Folder:               cloudflare.F(email_security.SettingDomainBatchParamsPatchesFolderAllItems),
			IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			LookbackHops:         cloudflare.F(int64(1)),
			Regions:              cloudflare.F([]email_security.SettingDomainBatchParamsPatchesRegion{email_security.SettingDomainBatchParamsPatchesRegionGlobal}),
			RequireTLSInbound:    cloudflare.F(true),
			RequireTLSOutbound:   cloudflare.F(true),
			Transport:            cloudflare.F("transport"),
		}}),
		Posts: cloudflare.F([]email_security.SettingDomainBatchParamsPost{{
			AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainBatchParamsPostsAllowedDeliveryMode{email_security.SettingDomainBatchParamsPostsAllowedDeliveryModeDirect}),
			Domain:               cloudflare.F("domain"),
			DropDispositions:     cloudflare.F([]email_security.SettingDomainBatchParamsPostsDropDisposition{email_security.SettingDomainBatchParamsPostsDropDispositionMalicious}),
			IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			Regions:              cloudflare.F([]email_security.SettingDomainBatchParamsPostsRegion{email_security.SettingDomainBatchParamsPostsRegionGlobal}),
			Folder:               cloudflare.F(email_security.SettingDomainBatchParamsPostsFolderAllItems),
			IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LookbackHops:         cloudflare.F(int64(1)),
			RequireTLSInbound:    cloudflare.F(true),
			RequireTLSOutbound:   cloudflare.F(true),
			Transport:            cloudflare.F("transport"),
		}}),
		Puts: cloudflare.F([]email_security.SettingDomainBatchParamsPut{{
			ID:                   cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
			AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainBatchParamsPutsAllowedDeliveryMode{email_security.SettingDomainBatchParamsPutsAllowedDeliveryModeDirect}),
			DropDispositions:     cloudflare.F([]email_security.SettingDomainBatchParamsPutsDropDisposition{email_security.SettingDomainBatchParamsPutsDropDispositionMalicious}),
			IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			Regions:              cloudflare.F([]email_security.SettingDomainBatchParamsPutsRegion{email_security.SettingDomainBatchParamsPutsRegionGlobal}),
			Folder:               cloudflare.F(email_security.SettingDomainBatchParamsPutsFolderAllItems),
			IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LookbackHops:         cloudflare.F(int64(1)),
			RequireTLSInbound:    cloudflare.F(true),
			RequireTLSOutbound:   cloudflare.F(true),
			Transport:            cloudflare.F("transport"),
		}}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingDomainBulkDelete(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.BulkDelete(context.TODO(), email_security.SettingDomainBulkDeleteParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettingDomainEditWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.Edit(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingDomainEditParams{
			AccountID:            cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedDeliveryModes: cloudflare.F([]email_security.SettingDomainEditParamsAllowedDeliveryMode{email_security.SettingDomainEditParamsAllowedDeliveryModeDirect}),
			DropDispositions:     cloudflare.F([]email_security.SettingDomainEditParamsDropDisposition{email_security.SettingDomainEditParamsDropDispositionMalicious}),
			Folder:               cloudflare.F(email_security.SettingDomainEditParamsFolderAllItems),
			IntegrationID:        cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IPRestrictions:       cloudflare.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			LookbackHops:         cloudflare.F(int64(1)),
			Regions:              cloudflare.F([]email_security.SettingDomainEditParamsRegion{email_security.SettingDomainEditParamsRegionGlobal}),
			RequireTLSInbound:    cloudflare.F(true),
			RequireTLSOutbound:   cloudflare.F(true),
			Transport:            cloudflare.F("transport"),
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

func TestSettingDomainGet(t *testing.T) {
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
	_, err := client.EmailSecurity.Settings.Domains.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		email_security.SettingDomainGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneEmailRoutingEnableService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneEmailRoutingEnableService]
// method instead.
type ZoneEmailRoutingEnableService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingEnableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingEnableService(opts ...option.RequestOption) (r *ZoneEmailRoutingEnableService) {
	r = &ZoneEmailRoutingEnableService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *ZoneEmailRoutingEnableService) EmailRoutingSettingsEnableEmailRouting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse struct {
	Errors   []ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON    `json:"-"`
}

// zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse]
type zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError]
type zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage]
type zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult struct {
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus `json:"status"`
	// Email Routing settings identifier.
	Tag  string                                                                         `json:"tag"`
	JSON zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON `json:"-"`
}

// zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult]
type zoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON struct {
	Created     apijson.Field
	Enabled     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	SkipWizard  apijson.Field
	Status      apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled bool

const (
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabledTrue  ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled = true
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabledFalse ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard bool

const (
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizardTrue  ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard = true
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizardFalse ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus string

const (
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusReady               ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "ready"
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusUnconfigured        ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "unconfigured"
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusMisconfigured       ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "misconfigured"
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusMisconfiguredLocked ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "misconfigured/locked"
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusUnlocked            ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess bool

const (
	ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccessTrue ZoneEmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess = true
)

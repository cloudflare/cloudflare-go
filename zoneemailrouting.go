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

// ZoneEmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneEmailRoutingService] method
// instead.
type ZoneEmailRoutingService struct {
	Options  []option.RequestOption
	Disables *ZoneEmailRoutingDisableService
	DNS      *ZoneEmailRoutingDNSService
	Enables  *ZoneEmailRoutingEnableService
	Rules    *ZoneEmailRoutingRuleService
}

// NewZoneEmailRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingService(opts ...option.RequestOption) (r *ZoneEmailRoutingService) {
	r = &ZoneEmailRoutingService{}
	r.Options = opts
	r.Disables = NewZoneEmailRoutingDisableService(opts...)
	r.DNS = NewZoneEmailRoutingDNSService(opts...)
	r.Enables = NewZoneEmailRoutingEnableService(opts...)
	r.Rules = NewZoneEmailRoutingRuleService(opts...)
	return
}

// Get information about the settings for your Email Routing zone.
func (r *ZoneEmailRoutingService) EmailRoutingSettingsGetEmailRoutingSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse struct {
	Errors   []ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON    `json:"-"`
}

// zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON contains
// the JSON metadata for the struct
// [ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse]
type zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError]
type zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage]
type zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult struct {
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus `json:"status"`
	// Email Routing settings identifier.
	Tag  string                                                                        `json:"tag"`
	JSON zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON `json:"-"`
}

// zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult]
type zoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON struct {
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

func (r *ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled bool

const (
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabledTrue  ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled = true
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabledFalse ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard bool

const (
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizardTrue  ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard = true
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizardFalse ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus string

const (
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusReady               ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "ready"
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusUnconfigured        ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "unconfigured"
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusMisconfigured       ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "misconfigured"
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusMisconfiguredLocked ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "misconfigured/locked"
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusUnlocked            ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess bool

const (
	ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccessTrue ZoneEmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess = true
)

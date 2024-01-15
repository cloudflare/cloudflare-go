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

// ZoneEmailRoutingDisableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneEmailRoutingDisableService] method instead.
type ZoneEmailRoutingDisableService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingDisableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingDisableService(opts ...option.RequestOption) (r *ZoneEmailRoutingDisableService) {
	r = &ZoneEmailRoutingDisableService{}
	r.Options = opts
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *ZoneEmailRoutingDisableService) EmailRoutingSettingsDisableEmailRouting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse struct {
	Errors   []ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON    `json:"-"`
}

// zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse]
type zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError]
type zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage]
type zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult struct {
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus `json:"status"`
	// Email Routing settings identifier.
	Tag  string                                                                           `json:"tag"`
	JSON zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON `json:"-"`
}

// zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult]
type zoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON struct {
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

func (r *ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled bool

const (
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabledTrue  ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled = true
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabledFalse ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard bool

const (
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizardTrue  ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard = true
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizardFalse ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus string

const (
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusReady               ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "ready"
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusUnconfigured        ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "unconfigured"
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusMisconfigured       ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "misconfigured"
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusMisconfiguredLocked ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "misconfigured/locked"
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusUnlocked            ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess bool

const (
	ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccessTrue ZoneEmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess = true
)

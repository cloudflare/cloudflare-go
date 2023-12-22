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
func (r *ZoneEmailRoutingService) EmailRoutingSettingsGetEmailRoutingSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmailSettingsResponseSingle struct {
	Errors   []EmailSettingsResponseSingleError   `json:"errors"`
	Messages []EmailSettingsResponseSingleMessage `json:"messages"`
	Result   EmailSettingsResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success EmailSettingsResponseSingleSuccess `json:"success"`
	JSON    emailSettingsResponseSingleJSON    `json:"-"`
}

// emailSettingsResponseSingleJSON contains the JSON metadata for the struct
// [EmailSettingsResponseSingle]
type emailSettingsResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSettingsResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSettingsResponseSingleError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    emailSettingsResponseSingleErrorJSON `json:"-"`
}

// emailSettingsResponseSingleErrorJSON contains the JSON metadata for the struct
// [EmailSettingsResponseSingleError]
type emailSettingsResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSettingsResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSettingsResponseSingleMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    emailSettingsResponseSingleMessageJSON `json:"-"`
}

// emailSettingsResponseSingleMessageJSON contains the JSON metadata for the struct
// [EmailSettingsResponseSingleMessage]
type emailSettingsResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSettingsResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSettingsResponseSingleResult struct {
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailSettingsResponseSingleResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailSettingsResponseSingleResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailSettingsResponseSingleResultStatus `json:"status"`
	// Email Routing settings identifier.
	Tag  string                                `json:"tag"`
	JSON emailSettingsResponseSingleResultJSON `json:"-"`
}

// emailSettingsResponseSingleResultJSON contains the JSON metadata for the struct
// [EmailSettingsResponseSingleResult]
type emailSettingsResponseSingleResultJSON struct {
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

func (r *EmailSettingsResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailSettingsResponseSingleResultEnabled bool

const (
	EmailSettingsResponseSingleResultEnabledTrue  EmailSettingsResponseSingleResultEnabled = true
	EmailSettingsResponseSingleResultEnabledFalse EmailSettingsResponseSingleResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailSettingsResponseSingleResultSkipWizard bool

const (
	EmailSettingsResponseSingleResultSkipWizardTrue  EmailSettingsResponseSingleResultSkipWizard = true
	EmailSettingsResponseSingleResultSkipWizardFalse EmailSettingsResponseSingleResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailSettingsResponseSingleResultStatus string

const (
	EmailSettingsResponseSingleResultStatusReady               EmailSettingsResponseSingleResultStatus = "ready"
	EmailSettingsResponseSingleResultStatusUnconfigured        EmailSettingsResponseSingleResultStatus = "unconfigured"
	EmailSettingsResponseSingleResultStatusMisconfigured       EmailSettingsResponseSingleResultStatus = "misconfigured"
	EmailSettingsResponseSingleResultStatusMisconfiguredLocked EmailSettingsResponseSingleResultStatus = "misconfigured/locked"
	EmailSettingsResponseSingleResultStatusUnlocked            EmailSettingsResponseSingleResultStatus = "unlocked"
)

// Whether the API call was successful
type EmailSettingsResponseSingleSuccess bool

const (
	EmailSettingsResponseSingleSuccessTrue EmailSettingsResponseSingleSuccess = true
)

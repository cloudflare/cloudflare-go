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

// EmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingService] method
// instead.
type EmailRoutingService struct {
	Options   []option.RequestOption
	Disables  *EmailRoutingDisableService
	DNS       *EmailRoutingDNSService
	Enables   *EmailRoutingEnableService
	Rules     *EmailRoutingRuleService
	Addresses *EmailRoutingAddressService
}

// NewEmailRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingService(opts ...option.RequestOption) (r *EmailRoutingService) {
	r = &EmailRoutingService{}
	r.Options = opts
	r.Disables = NewEmailRoutingDisableService(opts...)
	r.DNS = NewEmailRoutingDNSService(opts...)
	r.Enables = NewEmailRoutingEnableService(opts...)
	r.Rules = NewEmailRoutingRuleService(opts...)
	r.Addresses = NewEmailRoutingAddressService(opts...)
	return
}

// Get information about the settings for your Email Routing zone.
func (r *EmailRoutingService) EmailRoutingSettingsGetEmailRoutingSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse struct {
	Errors   []EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError   `json:"errors"`
	Messages []EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage `json:"messages"`
	Result   EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess `json:"success"`
	JSON    emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON    `json:"-"`
}

// emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON contains the
// JSON metadata for the struct
// [EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse]
type emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON `json:"-"`
}

// emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError]
type emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON `json:"-"`
}

// emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage]
type emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                                                    `json:"tag"`
	JSON emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON `json:"-"`
}

// emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult]
type emailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultJSON struct {
	ID          apijson.Field
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

func (r *EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled bool

const (
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabledTrue  EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled = true
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabledFalse EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard bool

const (
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizardTrue  EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard = true
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizardFalse EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus string

const (
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusReady               EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "ready"
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusUnconfigured        EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "unconfigured"
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusMisconfigured       EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "misconfigured"
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusMisconfiguredLocked EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "misconfigured/locked"
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatusUnlocked            EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess bool

const (
	EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccessTrue EmailRoutingEmailRoutingSettingsGetEmailRoutingSettingsResponseSuccess = true
)

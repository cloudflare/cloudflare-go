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

// EmailRoutingDisableService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingDisableService]
// method instead.
type EmailRoutingDisableService struct {
	Options []option.RequestOption
}

// NewEmailRoutingDisableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingDisableService(opts ...option.RequestOption) (r *EmailRoutingDisableService) {
	r = &EmailRoutingDisableService{}
	r.Options = opts
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *EmailRoutingDisableService) EmailRoutingSettingsDisableEmailRouting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse struct {
	Errors   []EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError   `json:"errors"`
	Messages []EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage `json:"messages"`
	Result   EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess `json:"success"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON    `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                                                       `json:"tag"`
	JSON emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultJSON struct {
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

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabledTrue  EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled = true
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabledFalse EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizardTrue  EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard = true
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizardFalse EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus string

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusReady               EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "ready"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusUnconfigured        EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "unconfigured"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusMisconfigured       EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "misconfigured"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusMisconfiguredLocked EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "misconfigured/locked"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatusUnlocked            EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccessTrue EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSuccess = true
)

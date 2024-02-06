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

// EmailRoutingEnableService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingEnableService] method
// instead.
type EmailRoutingEnableService struct {
	Options []option.RequestOption
}

// NewEmailRoutingEnableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingEnableService(opts ...option.RequestOption) (r *EmailRoutingEnableService) {
	r = &EmailRoutingEnableService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *EmailRoutingEnableService) EmailRoutingSettingsEnableEmailRouting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse struct {
	Errors   []EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError   `json:"errors"`
	Messages []EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage `json:"messages"`
	Result   EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess `json:"success"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON    `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                                                     `json:"tag"`
	JSON emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultJSON struct {
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

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabledTrue  EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled = true
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabledFalse EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizardTrue  EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard = true
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizardFalse EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus string

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusReady               EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "ready"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusUnconfigured        EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "unconfigured"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusMisconfigured       EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "misconfigured"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusMisconfiguredLocked EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "misconfigured/locked"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatusUnlocked            EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseResultStatus = "unlocked"
)

// Whether the API call was successful
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccessTrue EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSuccess = true
)

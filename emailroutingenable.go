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
	var env EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                                               `json:"tag"`
	JSON emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseJSON struct {
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

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabled bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabledTrue  EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabled = true
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabledFalse EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizard bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizardTrue  EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizard = true
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizardFalse EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus string

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatusReady               EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus = "ready"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatusUnconfigured        EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus = "unconfigured"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatusMisconfigured       EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus = "misconfigured"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatusMisconfiguredLocked EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus = "misconfigured/locked"
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatusUnlocked            EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseStatus = "unlocked"
)

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelope struct {
	Errors   []EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelope]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrors]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessages]
type emailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeSuccess bool

const (
	EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeSuccessTrue EmailRoutingEnableEmailRoutingSettingsEnableEmailRoutingResponseEnvelopeSuccess = true
)

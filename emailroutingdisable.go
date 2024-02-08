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
	var env EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                                                 `json:"tag"`
	JSON emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseJSON struct {
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

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabled bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabledTrue  EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabled = true
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabledFalse EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizard bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizardTrue  EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizard = true
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizardFalse EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus string

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatusReady               EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus = "ready"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatusUnconfigured        EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus = "unconfigured"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatusMisconfigured       EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus = "misconfigured"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatusMisconfiguredLocked EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus = "misconfigured/locked"
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatusUnlocked            EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseStatus = "unlocked"
)

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelope struct {
	Errors   []EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelope]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrors]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessages]
type emailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeSuccess bool

const (
	EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeSuccessTrue EmailRoutingDisableEmailRoutingSettingsDisableEmailRoutingResponseEnvelopeSuccess = true
)

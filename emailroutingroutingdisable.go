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

// EmailRoutingRoutingDisableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingDisableService] method instead.
type EmailRoutingRoutingDisableService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingDisableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingDisableService(opts ...option.RequestOption) (r *EmailRoutingRoutingDisableService) {
	r = &EmailRoutingRoutingDisableService{}
	r.Options = opts
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *EmailRoutingRoutingDisableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingDisableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingDisableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingDisableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingRoutingDisableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingRoutingDisableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingRoutingDisableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                    `json:"tag"`
	JSON emailRoutingRoutingDisableNewResponseJSON `json:"-"`
}

// emailRoutingRoutingDisableNewResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingDisableNewResponse]
type emailRoutingRoutingDisableNewResponseJSON struct {
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

func (r *EmailRoutingRoutingDisableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingRoutingDisableNewResponseEnabled bool

const (
	EmailRoutingRoutingDisableNewResponseEnabledTrue  EmailRoutingRoutingDisableNewResponseEnabled = true
	EmailRoutingRoutingDisableNewResponseEnabledFalse EmailRoutingRoutingDisableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingRoutingDisableNewResponseSkipWizard bool

const (
	EmailRoutingRoutingDisableNewResponseSkipWizardTrue  EmailRoutingRoutingDisableNewResponseSkipWizard = true
	EmailRoutingRoutingDisableNewResponseSkipWizardFalse EmailRoutingRoutingDisableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingRoutingDisableNewResponseStatus string

const (
	EmailRoutingRoutingDisableNewResponseStatusReady               EmailRoutingRoutingDisableNewResponseStatus = "ready"
	EmailRoutingRoutingDisableNewResponseStatusUnconfigured        EmailRoutingRoutingDisableNewResponseStatus = "unconfigured"
	EmailRoutingRoutingDisableNewResponseStatusMisconfigured       EmailRoutingRoutingDisableNewResponseStatus = "misconfigured"
	EmailRoutingRoutingDisableNewResponseStatusMisconfiguredLocked EmailRoutingRoutingDisableNewResponseStatus = "misconfigured/locked"
	EmailRoutingRoutingDisableNewResponseStatusUnlocked            EmailRoutingRoutingDisableNewResponseStatus = "unlocked"
)

type EmailRoutingRoutingDisableNewResponseEnvelope struct {
	Errors   []EmailRoutingRoutingDisableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingDisableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingDisableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingDisableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingDisableNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingDisableNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingDisableNewResponseEnvelope]
type emailRoutingRoutingDisableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDisableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingDisableNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingDisableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingDisableNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingDisableNewResponseEnvelopeErrors]
type emailRoutingRoutingDisableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDisableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingDisableNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingDisableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingDisableNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingDisableNewResponseEnvelopeMessages]
type emailRoutingRoutingDisableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDisableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingDisableNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingDisableNewResponseEnvelopeSuccessTrue EmailRoutingRoutingDisableNewResponseEnvelopeSuccess = true
)

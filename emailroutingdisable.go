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
func (r *EmailRoutingDisableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingDisableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingDisableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingDisableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingDisableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingDisableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingDisableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                             `json:"tag"`
	JSON emailRoutingDisableNewResponseJSON `json:"-"`
}

// emailRoutingDisableNewResponseJSON contains the JSON metadata for the struct
// [EmailRoutingDisableNewResponse]
type emailRoutingDisableNewResponseJSON struct {
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

func (r *EmailRoutingDisableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingDisableNewResponseEnabled bool

const (
	EmailRoutingDisableNewResponseEnabledTrue  EmailRoutingDisableNewResponseEnabled = true
	EmailRoutingDisableNewResponseEnabledFalse EmailRoutingDisableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingDisableNewResponseSkipWizard bool

const (
	EmailRoutingDisableNewResponseSkipWizardTrue  EmailRoutingDisableNewResponseSkipWizard = true
	EmailRoutingDisableNewResponseSkipWizardFalse EmailRoutingDisableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingDisableNewResponseStatus string

const (
	EmailRoutingDisableNewResponseStatusReady               EmailRoutingDisableNewResponseStatus = "ready"
	EmailRoutingDisableNewResponseStatusUnconfigured        EmailRoutingDisableNewResponseStatus = "unconfigured"
	EmailRoutingDisableNewResponseStatusMisconfigured       EmailRoutingDisableNewResponseStatus = "misconfigured"
	EmailRoutingDisableNewResponseStatusMisconfiguredLocked EmailRoutingDisableNewResponseStatus = "misconfigured/locked"
	EmailRoutingDisableNewResponseStatusUnlocked            EmailRoutingDisableNewResponseStatus = "unlocked"
)

type EmailRoutingDisableNewResponseEnvelope struct {
	Errors   []EmailRoutingDisableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingDisableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingDisableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingDisableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingDisableNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingDisableNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingDisableNewResponseEnvelope]
type emailRoutingDisableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    emailRoutingDisableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingDisableNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingDisableNewResponseEnvelopeErrors]
type emailRoutingDisableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDisableNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    emailRoutingDisableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingDisableNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingDisableNewResponseEnvelopeMessages]
type emailRoutingDisableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingDisableNewResponseEnvelopeSuccess bool

const (
	EmailRoutingDisableNewResponseEnvelopeSuccessTrue EmailRoutingDisableNewResponseEnvelopeSuccess = true
)

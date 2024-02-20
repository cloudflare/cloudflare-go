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
func (r *EmailRoutingEnableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingEnableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingEnableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingEnableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingEnableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingEnableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingEnableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                            `json:"tag"`
	JSON emailRoutingEnableNewResponseJSON `json:"-"`
}

// emailRoutingEnableNewResponseJSON contains the JSON metadata for the struct
// [EmailRoutingEnableNewResponse]
type emailRoutingEnableNewResponseJSON struct {
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

func (r *EmailRoutingEnableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingEnableNewResponseEnabled bool

const (
	EmailRoutingEnableNewResponseEnabledTrue  EmailRoutingEnableNewResponseEnabled = true
	EmailRoutingEnableNewResponseEnabledFalse EmailRoutingEnableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingEnableNewResponseSkipWizard bool

const (
	EmailRoutingEnableNewResponseSkipWizardTrue  EmailRoutingEnableNewResponseSkipWizard = true
	EmailRoutingEnableNewResponseSkipWizardFalse EmailRoutingEnableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingEnableNewResponseStatus string

const (
	EmailRoutingEnableNewResponseStatusReady               EmailRoutingEnableNewResponseStatus = "ready"
	EmailRoutingEnableNewResponseStatusUnconfigured        EmailRoutingEnableNewResponseStatus = "unconfigured"
	EmailRoutingEnableNewResponseStatusMisconfigured       EmailRoutingEnableNewResponseStatus = "misconfigured"
	EmailRoutingEnableNewResponseStatusMisconfiguredLocked EmailRoutingEnableNewResponseStatus = "misconfigured/locked"
	EmailRoutingEnableNewResponseStatusUnlocked            EmailRoutingEnableNewResponseStatus = "unlocked"
)

type EmailRoutingEnableNewResponseEnvelope struct {
	Errors   []EmailRoutingEnableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingEnableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingEnableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingEnableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingEnableNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingEnableNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingEnableNewResponseEnvelope]
type emailRoutingEnableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    emailRoutingEnableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingEnableNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingEnableNewResponseEnvelopeErrors]
type emailRoutingEnableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingEnableNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    emailRoutingEnableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingEnableNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [EmailRoutingEnableNewResponseEnvelopeMessages]
type emailRoutingEnableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingEnableNewResponseEnvelopeSuccess bool

const (
	EmailRoutingEnableNewResponseEnvelopeSuccessTrue EmailRoutingEnableNewResponseEnvelopeSuccess = true
)

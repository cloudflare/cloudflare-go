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

// EmailRoutingRoutingEnableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingEnableService] method instead.
type EmailRoutingRoutingEnableService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingEnableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingEnableService(opts ...option.RequestOption) (r *EmailRoutingRoutingEnableService) {
	r = &EmailRoutingRoutingEnableService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *EmailRoutingRoutingEnableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingEnableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingEnableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingEnableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingRoutingEnableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingRoutingEnableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingRoutingEnableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                                   `json:"tag"`
	JSON emailRoutingRoutingEnableNewResponseJSON `json:"-"`
}

// emailRoutingRoutingEnableNewResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingEnableNewResponse]
type emailRoutingRoutingEnableNewResponseJSON struct {
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

func (r *EmailRoutingRoutingEnableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingRoutingEnableNewResponseEnabled bool

const (
	EmailRoutingRoutingEnableNewResponseEnabledTrue  EmailRoutingRoutingEnableNewResponseEnabled = true
	EmailRoutingRoutingEnableNewResponseEnabledFalse EmailRoutingRoutingEnableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingRoutingEnableNewResponseSkipWizard bool

const (
	EmailRoutingRoutingEnableNewResponseSkipWizardTrue  EmailRoutingRoutingEnableNewResponseSkipWizard = true
	EmailRoutingRoutingEnableNewResponseSkipWizardFalse EmailRoutingRoutingEnableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingRoutingEnableNewResponseStatus string

const (
	EmailRoutingRoutingEnableNewResponseStatusReady               EmailRoutingRoutingEnableNewResponseStatus = "ready"
	EmailRoutingRoutingEnableNewResponseStatusUnconfigured        EmailRoutingRoutingEnableNewResponseStatus = "unconfigured"
	EmailRoutingRoutingEnableNewResponseStatusMisconfigured       EmailRoutingRoutingEnableNewResponseStatus = "misconfigured"
	EmailRoutingRoutingEnableNewResponseStatusMisconfiguredLocked EmailRoutingRoutingEnableNewResponseStatus = "misconfigured/locked"
	EmailRoutingRoutingEnableNewResponseStatusUnlocked            EmailRoutingRoutingEnableNewResponseStatus = "unlocked"
)

type EmailRoutingRoutingEnableNewResponseEnvelope struct {
	Errors   []EmailRoutingRoutingEnableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingEnableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingEnableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingEnableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingEnableNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingEnableNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingEnableNewResponseEnvelope]
type emailRoutingRoutingEnableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingEnableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingEnableNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    emailRoutingRoutingEnableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingEnableNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingEnableNewResponseEnvelopeErrors]
type emailRoutingRoutingEnableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingEnableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingEnableNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    emailRoutingRoutingEnableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingEnableNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingEnableNewResponseEnvelopeMessages]
type emailRoutingRoutingEnableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingEnableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingEnableNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingEnableNewResponseEnvelopeSuccessTrue EmailRoutingRoutingEnableNewResponseEnvelopeSuccess = true
)

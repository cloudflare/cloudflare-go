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
func (r *EmailRoutingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingGetResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingGetResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingGetResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingGetResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                      `json:"tag"`
	JSON emailRoutingGetResponseJSON `json:"-"`
}

// emailRoutingGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingGetResponse]
type emailRoutingGetResponseJSON struct {
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

func (r *EmailRoutingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// State of the zone settings for Email Routing.
type EmailRoutingGetResponseEnabled bool

const (
	EmailRoutingGetResponseEnabledTrue  EmailRoutingGetResponseEnabled = true
	EmailRoutingGetResponseEnabledFalse EmailRoutingGetResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingGetResponseSkipWizard bool

const (
	EmailRoutingGetResponseSkipWizardTrue  EmailRoutingGetResponseSkipWizard = true
	EmailRoutingGetResponseSkipWizardFalse EmailRoutingGetResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingGetResponseStatus string

const (
	EmailRoutingGetResponseStatusReady               EmailRoutingGetResponseStatus = "ready"
	EmailRoutingGetResponseStatusUnconfigured        EmailRoutingGetResponseStatus = "unconfigured"
	EmailRoutingGetResponseStatusMisconfigured       EmailRoutingGetResponseStatus = "misconfigured"
	EmailRoutingGetResponseStatusMisconfiguredLocked EmailRoutingGetResponseStatus = "misconfigured/locked"
	EmailRoutingGetResponseStatusUnlocked            EmailRoutingGetResponseStatus = "unlocked"
)

type EmailRoutingGetResponseEnvelope struct {
	Errors   []EmailRoutingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailRoutingGetResponseEnvelope]
type emailRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    emailRoutingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailRoutingGetResponseEnvelopeErrors]
type emailRoutingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    emailRoutingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [EmailRoutingGetResponseEnvelopeMessages]
type emailRoutingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingGetResponseEnvelopeSuccess bool

const (
	EmailRoutingGetResponseEnvelopeSuccessTrue EmailRoutingGetResponseEnvelopeSuccess = true
)

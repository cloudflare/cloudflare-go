// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// EmailRoutingRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingRoutingService]
// method instead.
type EmailRoutingRoutingService struct {
	Options   []option.RequestOption
	Disables  *EmailRoutingRoutingDisableService
	DNS       *EmailRoutingRoutingDNSService
	Enables   *EmailRoutingRoutingEnableService
	Rules     *EmailRoutingRoutingRuleService
	Addresses *EmailRoutingRoutingAddressService
}

// NewEmailRoutingRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingService(opts ...option.RequestOption) (r *EmailRoutingRoutingService) {
	r = &EmailRoutingRoutingService{}
	r.Options = opts
	r.Disables = NewEmailRoutingRoutingDisableService(opts...)
	r.DNS = NewEmailRoutingRoutingDNSService(opts...)
	r.Enables = NewEmailRoutingRoutingEnableService(opts...)
	r.Rules = NewEmailRoutingRoutingRuleService(opts...)
	r.Addresses = NewEmailRoutingRoutingAddressService(opts...)
	return
}

// Get information about the settings for your Email Routing zone.
func (r *EmailRoutingRoutingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingGetResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingRoutingGetResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingRoutingGetResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingRoutingGetResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                             `json:"tag"`
	JSON emailRoutingRoutingGetResponseJSON `json:"-"`
}

// emailRoutingRoutingGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRoutingGetResponse]
type emailRoutingRoutingGetResponseJSON struct {
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

func (r *EmailRoutingRoutingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingGetResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailRoutingRoutingGetResponseEnabled bool

const (
	EmailRoutingRoutingGetResponseEnabledTrue  EmailRoutingRoutingGetResponseEnabled = true
	EmailRoutingRoutingGetResponseEnabledFalse EmailRoutingRoutingGetResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingRoutingGetResponseSkipWizard bool

const (
	EmailRoutingRoutingGetResponseSkipWizardTrue  EmailRoutingRoutingGetResponseSkipWizard = true
	EmailRoutingRoutingGetResponseSkipWizardFalse EmailRoutingRoutingGetResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type EmailRoutingRoutingGetResponseStatus string

const (
	EmailRoutingRoutingGetResponseStatusReady               EmailRoutingRoutingGetResponseStatus = "ready"
	EmailRoutingRoutingGetResponseStatusUnconfigured        EmailRoutingRoutingGetResponseStatus = "unconfigured"
	EmailRoutingRoutingGetResponseStatusMisconfigured       EmailRoutingRoutingGetResponseStatus = "misconfigured"
	EmailRoutingRoutingGetResponseStatusMisconfiguredLocked EmailRoutingRoutingGetResponseStatus = "misconfigured/locked"
	EmailRoutingRoutingGetResponseStatusUnlocked            EmailRoutingRoutingGetResponseStatus = "unlocked"
)

type EmailRoutingRoutingGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingGetResponseEnvelope]
type emailRoutingRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    emailRoutingRoutingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingGetResponseEnvelopeErrors]
type emailRoutingRoutingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    emailRoutingRoutingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingGetResponseEnvelopeMessages]
type emailRoutingRoutingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingGetResponseEnvelopeSuccessTrue EmailRoutingRoutingGetResponseEnvelopeSuccess = true
)

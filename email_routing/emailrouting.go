// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingService] method
// instead.
type EmailRoutingService struct {
	Options   []option.RequestOption
	DNS       *DNSService
	Rules     *RuleService
	Addresses *AddressService
}

// NewEmailRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingService(opts ...option.RequestOption) (r *EmailRoutingService) {
	r = &EmailRoutingService{}
	r.Options = opts
	r.DNS = NewDNSService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Addresses = NewAddressService(opts...)
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *EmailRoutingService) Disable(ctx context.Context, zoneIdentifier string, body EmailRoutingDisableParams, opts ...option.RequestOption) (res *EmailSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingDisableResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *EmailRoutingService) Enable(ctx context.Context, zoneIdentifier string, body EmailRoutingEnableParams, opts ...option.RequestOption) (res *EmailSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingEnableResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about the settings for your Email Routing zone.
func (r *EmailRoutingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailSettings, err error) {
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

type EmailSettings struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailSettingsEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailSettingsSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailSettingsStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string            `json:"tag"`
	JSON emailSettingsJSON `json:"-"`
}

// emailSettingsJSON contains the JSON metadata for the struct [EmailSettings]
type emailSettingsJSON struct {
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

func (r *EmailSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSettingsJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailSettingsEnabled bool

const (
	EmailSettingsEnabledTrue  EmailSettingsEnabled = true
	EmailSettingsEnabledFalse EmailSettingsEnabled = false
)

func (r EmailSettingsEnabled) IsKnown() bool {
	switch r {
	case EmailSettingsEnabledTrue, EmailSettingsEnabledFalse:
		return true
	}
	return false
}

// Flag to check if the user skipped the configuration wizard.
type EmailSettingsSkipWizard bool

const (
	EmailSettingsSkipWizardTrue  EmailSettingsSkipWizard = true
	EmailSettingsSkipWizardFalse EmailSettingsSkipWizard = false
)

func (r EmailSettingsSkipWizard) IsKnown() bool {
	switch r {
	case EmailSettingsSkipWizardTrue, EmailSettingsSkipWizardFalse:
		return true
	}
	return false
}

// Show the state of your account, and the type or configuration error.
type EmailSettingsStatus string

const (
	EmailSettingsStatusReady               EmailSettingsStatus = "ready"
	EmailSettingsStatusUnconfigured        EmailSettingsStatus = "unconfigured"
	EmailSettingsStatusMisconfigured       EmailSettingsStatus = "misconfigured"
	EmailSettingsStatusMisconfiguredLocked EmailSettingsStatus = "misconfigured/locked"
	EmailSettingsStatusUnlocked            EmailSettingsStatus = "unlocked"
)

func (r EmailSettingsStatus) IsKnown() bool {
	switch r {
	case EmailSettingsStatusReady, EmailSettingsStatusUnconfigured, EmailSettingsStatusMisconfigured, EmailSettingsStatusMisconfiguredLocked, EmailSettingsStatusUnlocked:
		return true
	}
	return false
}

type EmailRoutingDisableParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r EmailRoutingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EmailRoutingDisableResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   EmailSettings                                             `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingDisableResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingDisableResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingDisableResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingDisableResponseEnvelope]
type emailRoutingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingDisableResponseEnvelopeSuccess bool

const (
	EmailRoutingDisableResponseEnvelopeSuccessTrue EmailRoutingDisableResponseEnvelopeSuccess = true
)

func (r EmailRoutingDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailRoutingEnableParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r EmailRoutingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EmailRoutingEnableResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   EmailSettings                                             `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingEnableResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingEnableResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingEnableResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailRoutingEnableResponseEnvelope]
type emailRoutingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingEnableResponseEnvelopeSuccess bool

const (
	EmailRoutingEnableResponseEnvelopeSuccessTrue EmailRoutingEnableResponseEnvelopeSuccess = true
)

func (r EmailRoutingEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailRoutingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   EmailSettings                                             `json:"result,required"`
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

func (r emailRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingGetResponseEnvelopeSuccess bool

const (
	EmailRoutingGetResponseEnvelopeSuccessTrue EmailRoutingGetResponseEnvelopeSuccess = true
)

func (r EmailRoutingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

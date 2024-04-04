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
func (r *EmailRoutingService) Disable(ctx context.Context, zoneIdentifier string, body EmailRoutingDisableParams, opts ...option.RequestOption) (res *EmailRoutingDisableResponse, err error) {
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
func (r *EmailRoutingService) Enable(ctx context.Context, zoneIdentifier string, body EmailRoutingEnableParams, opts ...option.RequestOption) (res *EmailRoutingEnableResponse, err error) {
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

type EmailRoutingDisableResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingDisableResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingDisableResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingDisableResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                          `json:"tag"`
	JSON emailRoutingDisableResponseJSON `json:"-"`
}

// emailRoutingDisableResponseJSON contains the JSON metadata for the struct
// [EmailRoutingDisableResponse]
type emailRoutingDisableResponseJSON struct {
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

func (r *EmailRoutingDisableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingDisableResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailRoutingDisableResponseEnabled bool

const (
	EmailRoutingDisableResponseEnabledTrue  EmailRoutingDisableResponseEnabled = true
	EmailRoutingDisableResponseEnabledFalse EmailRoutingDisableResponseEnabled = false
)

func (r EmailRoutingDisableResponseEnabled) IsKnown() bool {
	switch r {
	case EmailRoutingDisableResponseEnabledTrue, EmailRoutingDisableResponseEnabledFalse:
		return true
	}
	return false
}

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingDisableResponseSkipWizard bool

const (
	EmailRoutingDisableResponseSkipWizardTrue  EmailRoutingDisableResponseSkipWizard = true
	EmailRoutingDisableResponseSkipWizardFalse EmailRoutingDisableResponseSkipWizard = false
)

func (r EmailRoutingDisableResponseSkipWizard) IsKnown() bool {
	switch r {
	case EmailRoutingDisableResponseSkipWizardTrue, EmailRoutingDisableResponseSkipWizardFalse:
		return true
	}
	return false
}

// Show the state of your account, and the type or configuration error.
type EmailRoutingDisableResponseStatus string

const (
	EmailRoutingDisableResponseStatusReady               EmailRoutingDisableResponseStatus = "ready"
	EmailRoutingDisableResponseStatusUnconfigured        EmailRoutingDisableResponseStatus = "unconfigured"
	EmailRoutingDisableResponseStatusMisconfigured       EmailRoutingDisableResponseStatus = "misconfigured"
	EmailRoutingDisableResponseStatusMisconfiguredLocked EmailRoutingDisableResponseStatus = "misconfigured/locked"
	EmailRoutingDisableResponseStatusUnlocked            EmailRoutingDisableResponseStatus = "unlocked"
)

func (r EmailRoutingDisableResponseStatus) IsKnown() bool {
	switch r {
	case EmailRoutingDisableResponseStatusReady, EmailRoutingDisableResponseStatusUnconfigured, EmailRoutingDisableResponseStatusMisconfigured, EmailRoutingDisableResponseStatusMisconfiguredLocked, EmailRoutingDisableResponseStatusUnlocked:
		return true
	}
	return false
}

type EmailRoutingEnableResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled EmailRoutingEnableResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailRoutingEnableResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailRoutingEnableResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                         `json:"tag"`
	JSON emailRoutingEnableResponseJSON `json:"-"`
}

// emailRoutingEnableResponseJSON contains the JSON metadata for the struct
// [EmailRoutingEnableResponse]
type emailRoutingEnableResponseJSON struct {
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

func (r *EmailRoutingEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingEnableResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailRoutingEnableResponseEnabled bool

const (
	EmailRoutingEnableResponseEnabledTrue  EmailRoutingEnableResponseEnabled = true
	EmailRoutingEnableResponseEnabledFalse EmailRoutingEnableResponseEnabled = false
)

func (r EmailRoutingEnableResponseEnabled) IsKnown() bool {
	switch r {
	case EmailRoutingEnableResponseEnabledTrue, EmailRoutingEnableResponseEnabledFalse:
		return true
	}
	return false
}

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingEnableResponseSkipWizard bool

const (
	EmailRoutingEnableResponseSkipWizardTrue  EmailRoutingEnableResponseSkipWizard = true
	EmailRoutingEnableResponseSkipWizardFalse EmailRoutingEnableResponseSkipWizard = false
)

func (r EmailRoutingEnableResponseSkipWizard) IsKnown() bool {
	switch r {
	case EmailRoutingEnableResponseSkipWizardTrue, EmailRoutingEnableResponseSkipWizardFalse:
		return true
	}
	return false
}

// Show the state of your account, and the type or configuration error.
type EmailRoutingEnableResponseStatus string

const (
	EmailRoutingEnableResponseStatusReady               EmailRoutingEnableResponseStatus = "ready"
	EmailRoutingEnableResponseStatusUnconfigured        EmailRoutingEnableResponseStatus = "unconfigured"
	EmailRoutingEnableResponseStatusMisconfigured       EmailRoutingEnableResponseStatus = "misconfigured"
	EmailRoutingEnableResponseStatusMisconfiguredLocked EmailRoutingEnableResponseStatus = "misconfigured/locked"
	EmailRoutingEnableResponseStatusUnlocked            EmailRoutingEnableResponseStatus = "unlocked"
)

func (r EmailRoutingEnableResponseStatus) IsKnown() bool {
	switch r {
	case EmailRoutingEnableResponseStatusReady, EmailRoutingEnableResponseStatusUnconfigured, EmailRoutingEnableResponseStatusMisconfigured, EmailRoutingEnableResponseStatusMisconfiguredLocked, EmailRoutingEnableResponseStatusUnlocked:
		return true
	}
	return false
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

func (r emailRoutingGetResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailRoutingGetResponseEnabled bool

const (
	EmailRoutingGetResponseEnabledTrue  EmailRoutingGetResponseEnabled = true
	EmailRoutingGetResponseEnabledFalse EmailRoutingGetResponseEnabled = false
)

func (r EmailRoutingGetResponseEnabled) IsKnown() bool {
	switch r {
	case EmailRoutingGetResponseEnabledTrue, EmailRoutingGetResponseEnabledFalse:
		return true
	}
	return false
}

// Flag to check if the user skipped the configuration wizard.
type EmailRoutingGetResponseSkipWizard bool

const (
	EmailRoutingGetResponseSkipWizardTrue  EmailRoutingGetResponseSkipWizard = true
	EmailRoutingGetResponseSkipWizardFalse EmailRoutingGetResponseSkipWizard = false
)

func (r EmailRoutingGetResponseSkipWizard) IsKnown() bool {
	switch r {
	case EmailRoutingGetResponseSkipWizardTrue, EmailRoutingGetResponseSkipWizardFalse:
		return true
	}
	return false
}

// Show the state of your account, and the type or configuration error.
type EmailRoutingGetResponseStatus string

const (
	EmailRoutingGetResponseStatusReady               EmailRoutingGetResponseStatus = "ready"
	EmailRoutingGetResponseStatusUnconfigured        EmailRoutingGetResponseStatus = "unconfigured"
	EmailRoutingGetResponseStatusMisconfigured       EmailRoutingGetResponseStatus = "misconfigured"
	EmailRoutingGetResponseStatusMisconfiguredLocked EmailRoutingGetResponseStatus = "misconfigured/locked"
	EmailRoutingGetResponseStatusUnlocked            EmailRoutingGetResponseStatus = "unlocked"
)

func (r EmailRoutingGetResponseStatus) IsKnown() bool {
	switch r {
	case EmailRoutingGetResponseStatusReady, EmailRoutingGetResponseStatusUnconfigured, EmailRoutingGetResponseStatusMisconfigured, EmailRoutingGetResponseStatusMisconfiguredLocked, EmailRoutingGetResponseStatusUnlocked:
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   EmailRoutingDisableResponse  `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   EmailRoutingEnableResponse   `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   EmailRoutingGetResponse      `json:"result,required"`
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RoutingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRoutingService] method instead.
type RoutingService struct {
	Options   []option.RequestOption
	DNS       *RoutingDNSService
	Rules     *RoutingRuleService
	Addresses *RoutingAddressService
}

// NewRoutingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoutingService(opts ...option.RequestOption) (r *RoutingService) {
	r = &RoutingService{}
	r.Options = opts
	r.DNS = NewRoutingDNSService(opts...)
	r.Rules = NewRoutingRuleService(opts...)
	r.Addresses = NewRoutingAddressService(opts...)
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *RoutingService) Disable(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingDisableResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingDisableResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *RoutingService) Enable(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingEnableResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingEnableResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about the settings for your Email Routing zone.
func (r *RoutingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoutingDisableResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled RoutingDisableResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard RoutingDisableResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status RoutingDisableResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                     `json:"tag"`
	JSON routingDisableResponseJSON `json:"-"`
}

// routingDisableResponseJSON contains the JSON metadata for the struct
// [RoutingDisableResponse]
type routingDisableResponseJSON struct {
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

func (r *RoutingDisableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type RoutingDisableResponseEnabled bool

const (
	RoutingDisableResponseEnabledTrue  RoutingDisableResponseEnabled = true
	RoutingDisableResponseEnabledFalse RoutingDisableResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type RoutingDisableResponseSkipWizard bool

const (
	RoutingDisableResponseSkipWizardTrue  RoutingDisableResponseSkipWizard = true
	RoutingDisableResponseSkipWizardFalse RoutingDisableResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type RoutingDisableResponseStatus string

const (
	RoutingDisableResponseStatusReady               RoutingDisableResponseStatus = "ready"
	RoutingDisableResponseStatusUnconfigured        RoutingDisableResponseStatus = "unconfigured"
	RoutingDisableResponseStatusMisconfigured       RoutingDisableResponseStatus = "misconfigured"
	RoutingDisableResponseStatusMisconfiguredLocked RoutingDisableResponseStatus = "misconfigured/locked"
	RoutingDisableResponseStatusUnlocked            RoutingDisableResponseStatus = "unlocked"
)

type RoutingEnableResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled RoutingEnableResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard RoutingEnableResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status RoutingEnableResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                    `json:"tag"`
	JSON routingEnableResponseJSON `json:"-"`
}

// routingEnableResponseJSON contains the JSON metadata for the struct
// [RoutingEnableResponse]
type routingEnableResponseJSON struct {
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

func (r *RoutingEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type RoutingEnableResponseEnabled bool

const (
	RoutingEnableResponseEnabledTrue  RoutingEnableResponseEnabled = true
	RoutingEnableResponseEnabledFalse RoutingEnableResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type RoutingEnableResponseSkipWizard bool

const (
	RoutingEnableResponseSkipWizardTrue  RoutingEnableResponseSkipWizard = true
	RoutingEnableResponseSkipWizardFalse RoutingEnableResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type RoutingEnableResponseStatus string

const (
	RoutingEnableResponseStatusReady               RoutingEnableResponseStatus = "ready"
	RoutingEnableResponseStatusUnconfigured        RoutingEnableResponseStatus = "unconfigured"
	RoutingEnableResponseStatusMisconfigured       RoutingEnableResponseStatus = "misconfigured"
	RoutingEnableResponseStatusMisconfiguredLocked RoutingEnableResponseStatus = "misconfigured/locked"
	RoutingEnableResponseStatusUnlocked            RoutingEnableResponseStatus = "unlocked"
)

type RoutingGetResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled RoutingGetResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard RoutingGetResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status RoutingGetResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                 `json:"tag"`
	JSON routingGetResponseJSON `json:"-"`
}

// routingGetResponseJSON contains the JSON metadata for the struct
// [RoutingGetResponse]
type routingGetResponseJSON struct {
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

func (r *RoutingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingGetResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type RoutingGetResponseEnabled bool

const (
	RoutingGetResponseEnabledTrue  RoutingGetResponseEnabled = true
	RoutingGetResponseEnabledFalse RoutingGetResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type RoutingGetResponseSkipWizard bool

const (
	RoutingGetResponseSkipWizardTrue  RoutingGetResponseSkipWizard = true
	RoutingGetResponseSkipWizardFalse RoutingGetResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type RoutingGetResponseStatus string

const (
	RoutingGetResponseStatusReady               RoutingGetResponseStatus = "ready"
	RoutingGetResponseStatusUnconfigured        RoutingGetResponseStatus = "unconfigured"
	RoutingGetResponseStatusMisconfigured       RoutingGetResponseStatus = "misconfigured"
	RoutingGetResponseStatusMisconfiguredLocked RoutingGetResponseStatus = "misconfigured/locked"
	RoutingGetResponseStatusUnlocked            RoutingGetResponseStatus = "unlocked"
)

type RoutingDisableResponseEnvelope struct {
	Errors   []RoutingDisableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingDisableResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingDisableResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingDisableResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingDisableResponseEnvelopeJSON    `json:"-"`
}

// routingDisableResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingDisableResponseEnvelope]
type routingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingDisableResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    routingDisableResponseEnvelopeErrorsJSON `json:"-"`
}

// routingDisableResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingDisableResponseEnvelopeErrors]
type routingDisableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingDisableResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    routingDisableResponseEnvelopeMessagesJSON `json:"-"`
}

// routingDisableResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingDisableResponseEnvelopeMessages]
type routingDisableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingDisableResponseEnvelopeSuccess bool

const (
	RoutingDisableResponseEnvelopeSuccessTrue RoutingDisableResponseEnvelopeSuccess = true
)

type RoutingEnableResponseEnvelope struct {
	Errors   []RoutingEnableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingEnableResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingEnableResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingEnableResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingEnableResponseEnvelopeJSON    `json:"-"`
}

// routingEnableResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingEnableResponseEnvelope]
type routingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingEnableResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    routingEnableResponseEnvelopeErrorsJSON `json:"-"`
}

// routingEnableResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingEnableResponseEnvelopeErrors]
type routingEnableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingEnableResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    routingEnableResponseEnvelopeMessagesJSON `json:"-"`
}

// routingEnableResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingEnableResponseEnvelopeMessages]
type routingEnableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingEnableResponseEnvelopeSuccess bool

const (
	RoutingEnableResponseEnvelopeSuccessTrue RoutingEnableResponseEnvelopeSuccess = true
)

type RoutingGetResponseEnvelope struct {
	Errors   []RoutingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingGetResponseEnvelopeJSON    `json:"-"`
}

// routingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingGetResponseEnvelope]
type routingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    routingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RoutingGetResponseEnvelopeErrors]
type routingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    routingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RoutingGetResponseEnvelopeMessages]
type routingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingGetResponseEnvelopeSuccess bool

const (
	RoutingGetResponseEnvelopeSuccessTrue RoutingGetResponseEnvelopeSuccess = true
)

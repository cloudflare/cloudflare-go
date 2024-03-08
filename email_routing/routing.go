// File generated from our OpenAPI spec by Stainless.

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
	Disables  *RoutingDisableService
	DNS       *RoutingDNSService
	Enables   *RoutingEnableService
	Rules     *RoutingRuleService
	Addresses *RoutingAddressService
}

// NewRoutingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoutingService(opts ...option.RequestOption) (r *RoutingService) {
	r = &RoutingService{}
	r.Options = opts
	r.Disables = NewRoutingDisableService(opts...)
	r.DNS = NewRoutingDNSService(opts...)
	r.Enables = NewRoutingEnableService(opts...)
	r.Rules = NewRoutingRuleService(opts...)
	r.Addresses = NewRoutingAddressService(opts...)
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

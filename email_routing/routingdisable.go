// File generated from our OpenAPI spec by Stainless.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RoutingDisableService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingDisableService] method
// instead.
type RoutingDisableService struct {
	Options []option.RequestOption
}

// NewRoutingDisableService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingDisableService(opts ...option.RequestOption) (r *RoutingDisableService) {
	r = &RoutingDisableService{}
	r.Options = opts
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *RoutingDisableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingDisableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingDisableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoutingDisableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled RoutingDisableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard RoutingDisableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status RoutingDisableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                        `json:"tag"`
	JSON routingDisableNewResponseJSON `json:"-"`
}

// routingDisableNewResponseJSON contains the JSON metadata for the struct
// [RoutingDisableNewResponse]
type routingDisableNewResponseJSON struct {
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

func (r *RoutingDisableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableNewResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type RoutingDisableNewResponseEnabled bool

const (
	RoutingDisableNewResponseEnabledTrue  RoutingDisableNewResponseEnabled = true
	RoutingDisableNewResponseEnabledFalse RoutingDisableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type RoutingDisableNewResponseSkipWizard bool

const (
	RoutingDisableNewResponseSkipWizardTrue  RoutingDisableNewResponseSkipWizard = true
	RoutingDisableNewResponseSkipWizardFalse RoutingDisableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type RoutingDisableNewResponseStatus string

const (
	RoutingDisableNewResponseStatusReady               RoutingDisableNewResponseStatus = "ready"
	RoutingDisableNewResponseStatusUnconfigured        RoutingDisableNewResponseStatus = "unconfigured"
	RoutingDisableNewResponseStatusMisconfigured       RoutingDisableNewResponseStatus = "misconfigured"
	RoutingDisableNewResponseStatusMisconfiguredLocked RoutingDisableNewResponseStatus = "misconfigured/locked"
	RoutingDisableNewResponseStatusUnlocked            RoutingDisableNewResponseStatus = "unlocked"
)

type RoutingDisableNewResponseEnvelope struct {
	Errors   []RoutingDisableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingDisableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingDisableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingDisableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingDisableNewResponseEnvelopeJSON    `json:"-"`
}

// routingDisableNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingDisableNewResponseEnvelope]
type routingDisableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingDisableNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    routingDisableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// routingDisableNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingDisableNewResponseEnvelopeErrors]
type routingDisableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingDisableNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    routingDisableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// routingDisableNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingDisableNewResponseEnvelopeMessages]
type routingDisableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDisableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDisableNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingDisableNewResponseEnvelopeSuccess bool

const (
	RoutingDisableNewResponseEnvelopeSuccessTrue RoutingDisableNewResponseEnvelopeSuccess = true
)

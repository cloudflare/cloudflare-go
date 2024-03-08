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

// RoutingEnableService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingEnableService] method
// instead.
type RoutingEnableService struct {
	Options []option.RequestOption
}

// NewRoutingEnableService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingEnableService(opts ...option.RequestOption) (r *RoutingEnableService) {
	r = &RoutingEnableService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *RoutingEnableService) New(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RoutingEnableNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingEnableNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoutingEnableNewResponse struct {
	// Email Routing settings identifier.
	ID string `json:"id"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// State of the zone settings for Email Routing.
	Enabled RoutingEnableNewResponseEnabled `json:"enabled"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Domain of your zone.
	Name string `json:"name"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard RoutingEnableNewResponseSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status RoutingEnableNewResponseStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	Tag  string                       `json:"tag"`
	JSON routingEnableNewResponseJSON `json:"-"`
}

// routingEnableNewResponseJSON contains the JSON metadata for the struct
// [RoutingEnableNewResponse]
type routingEnableNewResponseJSON struct {
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

func (r *RoutingEnableNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableNewResponseJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type RoutingEnableNewResponseEnabled bool

const (
	RoutingEnableNewResponseEnabledTrue  RoutingEnableNewResponseEnabled = true
	RoutingEnableNewResponseEnabledFalse RoutingEnableNewResponseEnabled = false
)

// Flag to check if the user skipped the configuration wizard.
type RoutingEnableNewResponseSkipWizard bool

const (
	RoutingEnableNewResponseSkipWizardTrue  RoutingEnableNewResponseSkipWizard = true
	RoutingEnableNewResponseSkipWizardFalse RoutingEnableNewResponseSkipWizard = false
)

// Show the state of your account, and the type or configuration error.
type RoutingEnableNewResponseStatus string

const (
	RoutingEnableNewResponseStatusReady               RoutingEnableNewResponseStatus = "ready"
	RoutingEnableNewResponseStatusUnconfigured        RoutingEnableNewResponseStatus = "unconfigured"
	RoutingEnableNewResponseStatusMisconfigured       RoutingEnableNewResponseStatus = "misconfigured"
	RoutingEnableNewResponseStatusMisconfiguredLocked RoutingEnableNewResponseStatus = "misconfigured/locked"
	RoutingEnableNewResponseStatusUnlocked            RoutingEnableNewResponseStatus = "unlocked"
)

type RoutingEnableNewResponseEnvelope struct {
	Errors   []RoutingEnableNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingEnableNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingEnableNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingEnableNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingEnableNewResponseEnvelopeJSON    `json:"-"`
}

// routingEnableNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingEnableNewResponseEnvelope]
type routingEnableNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingEnableNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    routingEnableNewResponseEnvelopeErrorsJSON `json:"-"`
}

// routingEnableNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingEnableNewResponseEnvelopeErrors]
type routingEnableNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingEnableNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    routingEnableNewResponseEnvelopeMessagesJSON `json:"-"`
}

// routingEnableNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingEnableNewResponseEnvelopeMessages]
type routingEnableNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingEnableNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingEnableNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingEnableNewResponseEnvelopeSuccess bool

const (
	RoutingEnableNewResponseEnvelopeSuccessTrue RoutingEnableNewResponseEnvelopeSuccess = true
)

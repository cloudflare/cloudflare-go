// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// CTAlertingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCTAlertingService] method instead.
type CTAlertingService struct {
	Options []option.RequestOption
}

// NewCTAlertingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCTAlertingService(opts ...option.RequestOption) (r *CTAlertingService) {
	r = &CTAlertingService{}
	r.Options = opts
	return
}

// Create or update the Certificate Transparency alerting subscription for a zone.
// Enables or disables email notifications when certificates are issued for the
// zone's domains. For Free and Pro zones, the subscription is toggled on or off
// using the enabled field. Notification emails are sent to all users with SSL
// permissions on the zone. For Business and Enterprise zones, the emails field is
// required and controls which addresses receive alerts. Setting emails to an empty
// list disables the subscription regardless of the enabled field. A maximum of 10
// email addresses may be configured.
func (r *CTAlertingService) Edit(ctx context.Context, params CTAlertingEditParams, opts ...option.RequestOption) (res *CTAlertingEditResponse, err error) {
	var env CTAlertingEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ct/alerting", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the Certificate Transparency alerting subscription settings for a zone.
// Returns whether CT monitoring is enabled and, for Business and Enterprise zones,
// the list of email addresses that receive alerts.
func (r *CTAlertingService) Get(ctx context.Context, query CTAlertingGetParams, opts ...option.RequestOption) (res *CTAlertingGetResponse, err error) {
	var env CTAlertingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ct/alerting", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Certificate Transparency alerting subscription settings for a zone.
type CTAlertingEditResponse struct {
	// Whether CT alerting is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Email addresses that receive CT alert notifications. Only present and
	// configurable for Business and Enterprise zones. Maximum of 10 addresses. For
	// Free and Pro zones, notifications are sent to all users with SSL permissions on
	// the zone.
	Emails []string                   `json:"emails" format:"email"`
	JSON   ctAlertingEditResponseJSON `json:"-"`
}

// ctAlertingEditResponseJSON contains the JSON metadata for the struct
// [CTAlertingEditResponse]
type ctAlertingEditResponseJSON struct {
	Enabled     apijson.Field
	Emails      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Transparency alerting subscription settings for a zone.
type CTAlertingGetResponse struct {
	// Whether CT alerting is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Email addresses that receive CT alert notifications. Only present and
	// configurable for Business and Enterprise zones. Maximum of 10 addresses. For
	// Free and Pro zones, notifications are sent to all users with SSL permissions on
	// the zone.
	Emails []string                  `json:"emails" format:"email"`
	JSON   ctAlertingGetResponseJSON `json:"-"`
}

// ctAlertingGetResponseJSON contains the JSON metadata for the struct
// [CTAlertingGetResponse]
type ctAlertingGetResponseJSON struct {
	Enabled     apijson.Field
	Emails      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseJSON) RawJSON() string {
	return r.raw
}

type CTAlertingEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether CT alerting is enabled for the zone.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Email addresses that receive CT alert notifications. Only present and
	// configurable for Business and Enterprise zones. Maximum of 10 addresses. For
	// Free and Pro zones, notifications are sent to all users with SSL permissions on
	// the zone.
	Emails param.Field[[]string] `json:"emails" format:"email"`
}

func (r CTAlertingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CTAlertingEditResponseEnvelope struct {
	Errors   []CTAlertingEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CTAlertingEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success CTAlertingEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// Certificate Transparency alerting subscription settings for a zone.
	Result CTAlertingEditResponse             `json:"result"`
	JSON   ctAlertingEditResponseEnvelopeJSON `json:"-"`
}

// ctAlertingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTAlertingEditResponseEnvelope]
type ctAlertingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTAlertingEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           CTAlertingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             ctAlertingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ctAlertingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CTAlertingEditResponseEnvelopeErrors]
type ctAlertingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CTAlertingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CTAlertingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    ctAlertingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ctAlertingEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CTAlertingEditResponseEnvelopeErrorsSource]
type ctAlertingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CTAlertingEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           CTAlertingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             ctAlertingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ctAlertingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CTAlertingEditResponseEnvelopeMessages]
type ctAlertingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CTAlertingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CTAlertingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    ctAlertingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ctAlertingEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CTAlertingEditResponseEnvelopeMessagesSource]
type ctAlertingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CTAlertingEditResponseEnvelopeSuccess bool

const (
	CTAlertingEditResponseEnvelopeSuccessTrue CTAlertingEditResponseEnvelopeSuccess = true
)

func (r CTAlertingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CTAlertingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CTAlertingGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type CTAlertingGetResponseEnvelope struct {
	Errors   []CTAlertingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CTAlertingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success CTAlertingGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Certificate Transparency alerting subscription settings for a zone.
	Result CTAlertingGetResponse             `json:"result"`
	JSON   ctAlertingGetResponseEnvelopeJSON `json:"-"`
}

// ctAlertingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTAlertingGetResponseEnvelope]
type ctAlertingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTAlertingGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           CTAlertingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             ctAlertingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// ctAlertingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CTAlertingGetResponseEnvelopeErrors]
type ctAlertingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CTAlertingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CTAlertingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    ctAlertingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ctAlertingGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CTAlertingGetResponseEnvelopeErrorsSource]
type ctAlertingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CTAlertingGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           CTAlertingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             ctAlertingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// ctAlertingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CTAlertingGetResponseEnvelopeMessages]
type ctAlertingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CTAlertingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CTAlertingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    ctAlertingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ctAlertingGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CTAlertingGetResponseEnvelopeMessagesSource]
type ctAlertingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAlertingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAlertingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CTAlertingGetResponseEnvelopeSuccess bool

const (
	CTAlertingGetResponseEnvelopeSuccessTrue CTAlertingGetResponseEnvelopeSuccess = true
)

func (r CTAlertingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CTAlertingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
